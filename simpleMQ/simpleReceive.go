/*
@Description：
@Author : gilbert
@Date : 2022/6/3 13:54
*/

// 简单模式消费者

package main

import "rabbitMQ/rabbitMQ"

func main() {
	// 1.创建实例
	mq := rabbitMQ.NewRabbitMQSimple("testQ")
	// 2.消费消息
	mq.ConsumeSimple()
}
