package golanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)

	channel <- "INPUT"
}

func OnlyOut(channel <-chan string) {
	data := <-channel

	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)

	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}
func GiveMeResponse(response chan string) {
	response <- "BLA_BLA_BLA"

	time.Sleep(1 * time.Second)
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)

	go GiveMeResponse(channel)

	data := <-channel

	fmt.Println(data)

	defer close(channel)
}

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
