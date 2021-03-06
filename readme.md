### RabbitMQ 急速入门 
#### 快速开始
创建一个本地项目rabbitMQ，在根目录下使用`go get`命令安装对应的第三方库
> go get -u github.com/streadway/amqp
#### 2.目录结构
``` 
$ tree
├── go.mod
├── go.sum
├── pubSubMQ
│   ├── pubSubPublish.go
│   ├── pubSubReceiveOne.go
│   └── pubSubReceiveTwo.go
├── rabbitMQ
│   └── rabbitmq.go
├── readme.md
├── simpleMQ
│   ├── simplePublish.go
│   └── simpleReceive.go
└── workMQ
    ├── workPublish.go
    ├── workReceiveOne.go
    └── workReceiveTwo.go
```
- simpleMQ：简单模式
  - ![img](https://tva1.sinaimg.cn/large/e6c9d24ely1h2x8ukrgv5j20b5034mx5.jpg)
  - 消息产生着将消息放入队列
  - 消息的消费者(consumer) 监听(while) 消息队列,如果队列中有消息,就消费掉,消息被拿走后,自动从队列中删除(隐患 消息可能没有被消费者正确处理,已经从队列中消失了,造成消息的丢失)应用场景:聊天(中间有一个过度的服务器;p端,c端)
- workMQ：工作模式
  - ![img](https://tva1.sinaimg.cn/large/e6c9d24ely1h2x8uvkk02j209t03p0sr.jpg)
  - 消息产生者将消息放入队列消费者可以有多个,消费者1,消费者2,同时监听同一个队列,消息被消费?C1 C2共同争抢当前的消息队列内容,谁先拿到谁负责消费消息(隐患,高并发情况下,默认会产生某一个消息被多个消费者共同使用,可以设置一个开关(syncronize,与同步锁的性能不一样) 保证一条消息只能被一个消费者使用)
- 应用场景:红包;大项目中的资源调度(任务分配系统不需知道哪一个任务执行系统在空闲,直接将任务扔到消息队列中,空闲的系统自动争抢)
- pubSubMQ：订阅模式
  - ![img](https://tva1.sinaimg.cn/large/e6c9d24ely1h2x8v5ag17j20b604kwel.jpg)
  - X代表交换机rabbitMQ内部组件,erlang 消息产生者是代码完成,代码的执行效率不高,消息产生者将消息放入交换机,交换机发布订阅把消息发送到所有消息队列中,对应消息队列的消费者拿到消息进行消费
  - 相关场景:邮件群发,群聊天,广播(广告)
- routingMQ：路由模式
  - ![img](https://tva1.sinaimg.cn/large/e6c9d24ely1h2x8venumpj20bv04d74f.jpg)
  - 消息生产者将消息发送给交换机按照路由判断,路由是字符串(info) 当前产生的消息携带路由字符(对象的方法),交换机根据路由的key,只能匹配上路由key对应的消息队列,对应的消费者才能消费消息;
  - 根据业务功能定义路由字符串
  - 从系统的代码逻辑中获取对应的功能字符串,将消息任务扔到对应的队列中业务场景:error 通知;EXCEPTION;错误通知的功能;传统意义的错误通知;客户通知;利用key路由,可以将程序中的错误封装成消息传入到消息队列中,开发者可以自定义消费者,实时接收错误;
- topicMQ：
  - ![img](https://tva1.sinaimg.cn/large/e6c9d24ely1h2x8t5ahmfj20dj04d74g.jpg)
  - 话题模式，一个消息被多个消费者获取，消息的目标queue可用BindingKey以通配符，（#：一个或多个词，*：一个词）的方式指定
  - 星号井号代表通配符
  - 星号代表多个单词,井号代表一个单词
  - 路由功能添加模糊匹配
  - 消息产生者产生消息,把消息交给交换机
  - 交换机根据key的规则模糊匹配到对应的队列,由队列的监听消费者接收消息消费
