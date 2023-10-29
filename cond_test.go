package golanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)
var group = sync.WaitGroup{}

func WaitCondition(value int) {
	defer group.Done()

	group.Add(1)

	cond.L.Lock()

	cond.Wait()

	fmt.Println("DONE:", value)

	cond.L.Unlock()
}

func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		go WaitCondition(i)
	}

	// tell the locker to go one-on-one with cond.Signal()
	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		time.Sleep(1 * time.Second)

	// 		cond.Signal()
	// 	}
	// }()

	// tell the locker to run all at once with cond.Broadcast()
	go func() {
		time.Sleep(1 * time.Second)
		cond.Broadcast()
	}()

	group.Wait()
}
