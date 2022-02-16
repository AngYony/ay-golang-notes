# Go 并发编程概述

包含的内容：

- 基本并发原语（传统并发原语）：Mutex、RWMutex、Waitgroup、Cond、Pool、Context，这些都属于标准库中的并发原语，位于sync包。
- 原子操作：原子操作是其它并发原语的基础。
- Channel
- 扩展并发原语：信号量、SingleFlight、循环栅栏、ErrGroup等，这些不属于标准库但应用广泛。
- 分布式并发原语：Leader选举、分布式互斥锁、分布式读写锁、分布式队列等，这些一般基于第三方（如etcd）来实现的。



选择相应的并发原语的原则：

任务编排用 Channel，共享资源保护用传统并发原语。不可万事皆用 Channel。



各并发原语适用场景：

- 共享资源。并发地读写共享资源，会出现数据竞争的问题，所以需要 Mutex、RWMutex 这样的并发原语来保护。
- 任务编排。需要 goroutine 按照一定的规律执行，而 goroutine 之间有相互等待或者依赖的顺序关系，常常使用 WaitGroup 或者 Channel 来实现。
- 消息传递。信息交流以及不同的 goroutine 之间的线程安全的数据交流，常常使用 Channel 来实现。



相关工具：

- 推荐一个工具 Chronos - A static race detector for the go language， https://github.com/amit-davidson/Chronos
- 检测并发访问共享资源是否有问题的工具： [race detector](https://blog.golang.org/race-detector)，它可以帮助我们自动发现程序有没有 data race 的问题。