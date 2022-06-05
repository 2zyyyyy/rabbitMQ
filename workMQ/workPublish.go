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
		mq.PublishSimple("简单模式生产第" + strconv.Itoa(i) + "条数据")
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}
