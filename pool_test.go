package golanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() any {
			return "NEW_DATA" // default value
		},
	}

	var group sync.WaitGroup

	pool.Put("Yudistira")
	pool.Put("Eka")
	pool.Put("Pratama")

	for i := 0; i < 10; i++ {
		group.Add(1)
		go func() {
			defer group.Done()
			data := pool.Get()

			time.Sleep(1 * time.Second) // simulate when data has not been completed

			fmt.Println(data)
			pool.Put(data)
		}()
	}

	group.Wait()

	fmt.Println("Pool is finished")
}
