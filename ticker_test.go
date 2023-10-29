package golanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
	}()

	for {
		select {
		case tick, ok := <-ticker.C:
			if !ok {
				fmt.Println("Ticker stopped.")
				return
			}
			fmt.Println(tick)
		}
	}
}

func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)

	for tick := range channel {
		fmt.Println(tick)
	}
}
