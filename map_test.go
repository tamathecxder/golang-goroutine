package golanggoroutine

import (
	"fmt"
	"sync"
	"testing"
)

func AddToMap(data *sync.Map, group *sync.WaitGroup, value int) {
	data.Store(value, value)

	group.Add(1)
	defer group.Done()
}

func TestMap(t *testing.T) {
	data := &sync.Map{}
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go AddToMap(data, group, i)
	}

	group.Wait()

	data.Range(func(key, value any) bool {
		fmt.Println(key, ":", value)

		return true
	})
}
