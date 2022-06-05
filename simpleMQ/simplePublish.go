/*
@Description：
@Author : gilbert
@Date : 2022/6/3 13:53
*/

// 简单模式 生产者

package main

import (
	"fmt"
	"rabbitMQ/rabbitMQ"
)

func main() {
	// 1.创建实例
	mq := rabbitMQ.NewRabbitMQSimple("testQ")
	// 2.生产消息
	mq.PublishSimple("wanli test simple mq message.")
	fmt.Println("消息发送成功")
}
