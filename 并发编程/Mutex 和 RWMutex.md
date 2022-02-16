# Mutex 和 RWMutex

Mutex 用于解决多个 goroutine 并发更新同一个资源问题， 也被称为互斥锁或排它锁，作用类似于C# 中的lock。



互斥锁在实现时，牵扯到一个概念叫**临界区**。

临界区：多个 goroutine 都会访问或修改的资源。

一旦多个goroutine同步访问临界区，就会造成访问或操作错误，出现资源竞争的问题，所以需要使用互斥锁，限定临界区只能同时由一个goroutine持有。



## Mutex

简单来说，互斥锁 Mutex 就提供两个方法：Lock 和 Unlock，进入临界区之前调用 Lock 方法，退出临界区的时候调用 Unlock 方法。

```go
func(m *Mutex)Lock()
func(m *Mutex)Unlock()
```

具体表现为：

**当一个 goroutine 通过调用 Lock 方法，获得了这个锁的拥有权后，其它请求锁的 goroutine 就会阻塞在 Lock 方法的调用上，直到 锁被释放并且自己获取到了这个锁的拥有权。**

未使用Mutex，存在竞争问题的代码：

```go
func main() {
	var count = 0
	// 使用 WaitGroup 等待10个goroutine完成
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				// 注意：count++不是原子操作
				count++
			}
		}()
	}
	// 等待10个goroutine完成
	wg.Wait()
	fmt.Println(count)
}
```

上述代码需要特别注意：

count++ 不是一个原子操作，它至少包含几个步骤，比如读取变量 count 的当前值，对这个值加 1，把结果再保存到 count 中。因为不是原子操作，就可能有并发的问题。

可以通过 race detector 工具发现计数器程序的问题以及修复方法。

```
go run -race .\main.go
```

使用 Mutex 修复之后的代码：

```go
func main() {
	// 互斥锁保护计数器
	var mu sync.Mutex

	var count = 0
	// 使用 WaitGroup 等待10个goroutine完成
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				mu.Lock()
				// 注意：count++不是原子操作
				count++
				mu.Unlock()
			}
		}()
	}
	// 等待10个goroutine完成
	wg.Wait()
	fmt.Println(count)
}
```

注意：Mutex 的零值是还没有 goroutine 等待的未加锁的状态，所以你不需要额外的初始化，直接声明变量（如 var mu sync.Mutex）即可。



### Mutex 嵌入到 struct

最佳实践：

```go
type Counter struct {
	Name string

	mu    sync.Mutex
	count uint64
}

// 加1的方法，内部使用互斥锁保护
func (c *Counter) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

// 得到计算器的值，也需要锁保护
func (c *Counter) Count() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {
    // 封装好的计数器
	var counter Counter // 不需要显示初始化Mutex即可使用

    // 使用 WaitGroup 等待10个goroutine完成
	var wg sync.WaitGroup
	wg.Add(10)

    // 启动10个goroutine
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				counter.Incr() // 受到锁保护的方法
			}
		}()
	}
	// 等待10个goroutine完成
	wg.Wait()
	fmt.Println(counter.Count())
}
```

重点：

- 初始化嵌入的 struct 时，不需要显式初始化 Mutex 字段。Mutex不会因为没有初始化而出现空指针或者无法获取锁的情况。（上述代码第24行）
- 可以直接将 Mutex 作为字段嵌入，而无须指定字段名，这样就可以在struct上直接调用Lock/Unlock方法了。（上述代码第4行）
- 如果嵌入的 struct 有多个字段，一般会把 Mutex 放在要控制的字段上面，然后使用空行把字段分隔开来。（上述代码第3行）
- 通常建议将获取锁、释放锁、计数等逻辑封装成方法，而不对外暴露锁的逻辑。



等待的goroutine们是以FIFO排队的
1）当Mutex处于正常模式时，若此时没有新goroutine与队头goroutine竞争，则队头goroutine获得。若有新goroutine竞争大概率新goroutine获得。
2）当队头goroutine竞争锁失败1ms后，它会将Mutex调整为饥饿模式。进入饥饿模式后，锁的所有权会直接从解锁goroutine移交给队头goroutine，此时新来的goroutine直接放入队尾。

3）当一个goroutine获取锁后，如果发现自己满足下列条件中的任何一个#1它是队列中最后一个#2它等待锁的时间少于1ms，则将锁切换回正常模式











## RWMutex









