package golanggoroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	var counter int = 0

	for {
		select {
		case data := <-channel1:
			fmt.Println("Data from channel 1:", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data from channel 2:", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Value of " + strconv.Itoa(i+1)
		}

		close(channel)
	}()

	for data := range channel {
		fmt.Println("Receive:", data)
	}
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)

	defer close(channel)

	go func() {
		channel <- "Yudistira"
		channel <- "Eka"
		channel <- "Pratama"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	// channel <- "Anjim" // deadlock

	fmt.Println(cap(channel))
	fmt.Println(len(channel))

	time.Sleep(2 * time.Second)

	fmt.Println("DONE!")
}

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
func GiveMeResponse(channel chan string) {
	channel <- "BLA_BLA_BLA"

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
