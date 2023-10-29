package golanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)

	fmt.Println(time.Now())

	time := <-timer.C

	fmt.Println(time)
}

func TestTimeAfter(t *testing.T) {
	channel := time.After(1 * time.Second)

	fmt.Println(time.Now())

	time := <-channel

	fmt.Println(time)
}

func TestTimeAfterFunc(t *testing.T) {
	group := sync.WaitGroup{}

	group.Add(1)

	time.AfterFunc(3*time.Second, func() {
		fmt.Println("Executed after 3 seconds")

		group.Done()
	})

	group.Wait()
}
