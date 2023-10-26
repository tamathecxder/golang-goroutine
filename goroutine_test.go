package golanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

func RunGreetUhkti() {
	fmt.Println("Assalamualaikum, Ukhti")
}

func TestCreateGoroutine(t *testing.T) {
	go RunGreetUhkti()
	fmt.Println("OMG")

	time.Sleep(1 * time.Second)
}
