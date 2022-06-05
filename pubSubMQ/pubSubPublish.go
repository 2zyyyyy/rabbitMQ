/*
@Description：
@Author : gilbert
@Date : 2022/6/5 11:39
*/

package main

import (
	"fmt"
	"rabbitMQ/rabbitmq"
	"strconv"
	"time"
)

// 订阅模式生产者

func main() {
	mq := rabbitmq.NewRabbitMQPubSub("testEx")
	for i := 0; i < 100; i++ {
		mq.PublishPubSub("订阅模式生产第" + strconv.Itoa(i) + "条数据")
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}
