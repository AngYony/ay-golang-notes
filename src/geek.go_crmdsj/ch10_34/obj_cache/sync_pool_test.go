package obj_cache

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("创建了一个新对象")
			return 100
		},
	}

	v := pool.Get().(int) // 断言
	fmt.Println(v)
	pool.Put(3)
	runtime.GC() // GC 会清除sync.pool 中缓存的对象
	v1, _ := pool.Get().(int)
	fmt.Println(v1)

}

func TestSyncPoolInMultiGroutine(t *testing.T) {
	pool := &sync.Pool{
		// 找不到的时候调用New方法
		New: func() interface{} {
			fmt.Println("创建了一个新对象")
			return 10
		},
	}

	pool.Put(100)
	pool.Put(100)
	pool.Put(100)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			fmt.Println(pool.Get())
			wg.Done()
		}(i)
	}
	wg.Wait()
}
