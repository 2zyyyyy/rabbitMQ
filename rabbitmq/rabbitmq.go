/*
@Description：
@Author : gilbert
@Date : 2022/6/2 23:29
*/

package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
)

// url格式：amqp:账号:密码@rabbitmq 服务器地址:端口号/vhost

const (
	MQURL = "amqp://simpleU:123456@127.0.0.1:15672/simple"
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
func NewRabbitMQ(name, exchange, key string) *RabbitMQ {
	return &RabbitMQ{
		QueueName: name,
		Exchange:  exchange,
		Key:       key,
		MQURL:     MQURL,
	}
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
		panic(logInfo)
	}
}

// NewRabbitMQSimple simple 模式实例
func NewRabbitMQSimple(name string) *RabbitMQ {
	rabbitmq := NewRabbitMQ(name, "", "")
	// 创建 mq 连接
	var err error
	rabbitmq.conn, err = amqp.Dial(rabbitmq.MQURL)
	rabbitmq.failOnErr("创建连接失败", err)
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr("获取 channel 失败", err)
	return rabbitmq
}
