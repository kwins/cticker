package cticker

import (
	"fmt"
	"sync"
	"testing"
)

var wg sync.WaitGroup

func TestCTicker(t *testing.T) {
	ticker := NewTaskSchedule()
	for index := 0; index < 100; index++ {
		wg.Add(1)
		var task Task
		var i = index
		task.handler = func() error {
			fmt.Println("index:", i)
			wg.Done()
			return nil
		}
		ticker.AddTask(index+1, int64(index), &task)
	}
	ticker.CancelTask(10)
	wg.Wait()
}
