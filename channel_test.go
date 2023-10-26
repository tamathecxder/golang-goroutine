package golanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestRunChannel(t *testing.T) {
	channel := make(chan string)

	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)

		channel <- "SEND_DATA"

		fmt.Println("finish sending data to the channel")
	}()

	data := <-channel

	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)

	channel <- "anjim"

	// 1. directly assign into variable
	// data := <-channel

	// 2. set up the variables first and then assign them from the channel
	// var data string
	// data = <-channel

	// 3. receive and use the channel directly
	// fmt.Println(<-channel)

	defer close(channel)
}
