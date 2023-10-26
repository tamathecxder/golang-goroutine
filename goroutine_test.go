package golanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

func DisplayNumber(number int) {
	fmt.Println("Display:", number)
}

func RunGreetUhkti() {
	fmt.Println("Assalamualaikum, Ukhti")
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}

func TestCreateGoroutine(t *testing.T) {
	go RunGreetUhkti()
	fmt.Println("OMG")

	time.Sleep(1 * time.Second)
}
