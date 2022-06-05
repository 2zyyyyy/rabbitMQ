/*
@Description：
@Author : gilbert
@Date : 2022/6/2 23:29
*/

package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

// url格式：amqp:账号:密码@rabbitmq 服务器地址:端口号/vhost

const (
	MQURL = "amqp://simpleU:123456@localhost:5672/simple"
)

// RabbitMQ 结构体
type RabbitMQ struct {
	conn      *amqp.Connection
	channel   *amqp.Channel
	QueueName string // 队列名称
	Exchange  string // 交换机
	Key       string // key
	MQURL     string // 连接信息
}

// NewRabbitMQ 创建结构体实例
func NewRabbitMQ(queueName, exchange, key string) *RabbitMQ {
	rabbitmq := &RabbitMQ{
		QueueName: queueName,
		Exchange:  exchange,
		Key:       key,
		MQURL:     MQURL,
	}
	// 创建 mq 连接
	var err error
	rabbitmq.conn, err = amqp.Dial(rabbitmq.MQURL)
	rabbitmq.failOnErr("创建连接失败", err)
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr("获取 channel 失败", err)
	return rabbitmq
}

// Destroy 关闭通道和连接
func (r *RabbitMQ) Destroy() {
	err := r.channel.Close()
	r.failOnErr("channel 关闭失败", err)
	err = r.conn.Close()
	r.failOnErr("connections 关闭失败", err)
}

// failOnErr错误处理函数
func (r *RabbitMQ) failOnErr(msg string, err error) {
	if err != nil {
		logInfo := fmt.Sprintf("%s:%s", msg, err)
		fmt.Println(logInfo)
		panic(logInfo)
	}
}

// applyQueue 申请队列
func (r *RabbitMQ) applyQueue() {
	// 申请队列，如果不存在会自动创建 保证队列存在，消息可以发送到队列中
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		true,  // 是否持久化
		false, // 是否自动删除
		false, // 是否具有排他性
		false, // 是否阻塞
		nil,   // 额外属性
	)
	if err != nil {
		fmt.Println(err)
	}
}

// NewRabbitMQSimple simple 模式 step1:创建实例
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	return NewRabbitMQ(queueName, "", "")
}

// PublishSimple simple 模式 step2:生产消息
func (r *RabbitMQ) PublishSimple(msg string) {
	// 1.申请队列
	r.applyQueue()
	// 2.发送消息到队列中
	err := r.channel.Publish(
		r.Exchange,
		r.QueueName,
		false, // true:会根据 exchange 类型和 routkey 规则，当无法找到符合条件的队列会将该消息返还给发送者
		false, // true:当 exchange 发送到消息队列后发现队列上没有绑定消费者，会将消息返还给发送者
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
	if err != nil {
		fmt.Printf("发送消息到队列失败，err:%s\n", err)
		return
	}
}

// ConsumeSimple simple 模式 step3:消费消息
func (r *RabbitMQ) ConsumeSimple() {
	// 1.申请队列
	r.applyQueue()
	// 2.接收消息
	messages, err := r.channel.Consume(
		r.QueueName,
		"",    // 用来区分多个消费者
		true,  // 是否自动应答
		false, // 是否具有排他性
		false, // true:表示不能将同一个 connection 中发送的消息传递给这个 connection 中的消费者
		false, // true:表示消息需要阻塞
		nil,
	)
	if err != nil {
		fmt.Printf("接收消息失败,err:%s\n", err)
	}
	// 3.消费消息
	forever := make(chan bool) // 创建一个阻塞通道
	// 启用 goroutine 处理消息
	go func() {
		for msg := range messages {
			// 实现具体的消息处理逻辑，如落库等，这里简单打印
			log.Printf("received a message, 消费%s", msg.Body)
		}
	}()
	log.Printf("[*] Waiting for message, to exit ptrss CTRC+C")
	<-forever // 因为 forever 一直没有值 所以会一直等待 达到阻塞效果
}

// NewRabbitMQPubSub 订阅模式 step1:创建实例
func NewRabbitMQPubSub(exchangeName string) *RabbitMQ {
	// 1.创建实例
	mq := NewRabbitMQ("", exchangeName, "")
	var err error
	// 2.获取 connection
	mq.conn, err = amqp.Dial(mq.MQURL)
	mq.failOnErr("failed to connect mq", err)
	// 3.获取 channel
	mq.channel, err = mq.conn.Channel()
	mq.failOnErr("failed to open a channel", err)
	return mq
}

// PublishPubSub 订阅模式 step2:生产消息
func (r *RabbitMQ) PublishPubSub(message string) {
	// 1.尝试创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnErr("failed to declare an exchange", err)

	// 2.发送消息
	err = r.channel.Publish(
		r.Exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

// ConsumePubSub 订阅模式 step3:消费消息
func (r *RabbitMQ) ConsumePubSub() {
	// 1.试探性创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnErr("failed to declare an exchange", err)

	// 2.试探性创建队列，这里注意不写队列名称
	queue, err := r.channel.QueueDeclare(
		"", // 随机队列名称
		false,
		false,
		true,
		false,
		nil,
	)
	r.failOnErr("failed to declare a queue", err)

	// 3.绑定队列到 exchange 中
	err = r.channel.QueueBind(
		queue.Name,
		"", // 在订阅模式下,这里的 key 需要为空
		r.Exchange,
		false, nil,
	)

	// 4.消费消息
	messages, err := r.channel.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	forever := make(chan bool) // 创建一个阻塞通道
	// 启用 goroutine 处理消息
	go func() {
		for msg := range messages {
			// 实现具体的消息处理逻辑，如落库等，这里简单打印
			log.Printf("received a message, 消费%s", msg.Body)
		}
	}()
	log.Printf("[*] Waiting for message, to exit ptrss CTRC+C")
	<-forever // 因为 forever 一直没有值 所以会一直等待 达到阻塞效果
}
