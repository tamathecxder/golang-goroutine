package golanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchronous(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)

	fmt.Println("Test")

	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 1; i <= 100; i++ {
		go RunAsynchronous(group)
	}

	group.Wait()

	fmt.Println("WaitGroup is done!")
}
