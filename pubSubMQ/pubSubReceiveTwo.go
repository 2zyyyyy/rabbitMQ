/*
@Description：
@Author : gilbert
@Date : 2022/6/5 11:39
*/

package main

import "rabbitMQ/rabbitMQ"

// 订阅模式消费者 2

func main() {
	mq := rabbitMQ.NewRabbitMQPubSub("testEx")
	mq.ConsumePubSub()
}
