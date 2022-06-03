/*
@Description：
@Author : gilbert
@Date : 2022/6/3 23:18
*/

// work模式消费者1

package main

import "rabbitMQ/rabbitmq"

func main() {
	// 1.创建实例
	mq := rabbitmq.NewRabbitMQSimple("" + "testQ")
	// 2.消费消息
	mq.ConsumeSimple()
}
