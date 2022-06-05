/*
@Description：
@Author : gilbert
@Date : 2022/6/5 11:39
*/

package main

import "rabbitMQ/rabbitmq"

// 订阅模式消费者 1

func main() {
	mq := rabbitmq.NewRabbitMQPubSub("testEx")
	mq.ConsumePubSub()
}
