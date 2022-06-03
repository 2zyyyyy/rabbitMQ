/*
@Description：
@Author : gilbert
@Date : 2022/6/3 23:18
*/

package main

import (
	"fmt"
	"rabbitMQ/rabbitmq"
	"strconv"
	"time"
)

func main() {
	mq := rabbitmq.NewRabbitMQSimple("" + "testQ")
	for i := 0; i < 100; i++ {
		mq.PublishSimple("test work MQ message" + strconv.Itoa(i))
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}
