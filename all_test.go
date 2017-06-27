package cticker

import (
	"fmt"
	"testing"
	"time"
)

func TestCTicker(t *testing.T) {
	ticker := NewQueue(0, 0)
	for index := 0; index < 100; index++ {
		var task Task
		var i = index
		task.handler = func() error {
			fmt.Println("index:", i)
			return nil
		}
		ticker.AddTimerTask(index+1, fmt.Sprint(index), &task)
	}
	time.Sleep(time.Second * 5)
	ticker.CancelTask(fmt.Sprint(10))
	ta := ticker.GetTask(fmt.Sprint(15))
	ta.Cancel()
	time.Sleep(time.Hour)
}
