/*
@Description：
@Author : gilbert
@Date : 2022/6/3 23:18
*/

// work模式消费者2

package main

import "rabbitMQ/rabbitMQ"

func main() {
	// 1.创建实例
	mq := rabbitMQ.NewRabbitMQSimple("" + "testQ")
	// 2.消费消息
	mq.ConsumeSimple()
}
