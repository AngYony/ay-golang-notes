package share_mem

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++

		}()
	}
	time.Sleep(time.Second * 1)
	fmt.Printf("Counter= %d", counter)
}

func TestCounterThreadSafe(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer mut.Unlock()
			mut.Lock()
			counter++

		}()

	}

	time.Sleep(time.Second * 1)
	fmt.Println("counter = ", counter)
}

func TestCounterWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer mut.Unlock()
			mut.Lock()
			counter++

			wg.Done()
		}()

	}

	wg.Wait()
	fmt.Println("counter = ", counter)
}
