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
		task.handler = func() error {
			t.Log("index:", index)
			return nil
		}
		ticker.AddTimerTask(index+1, fmt.Sprint(index), &task)
	}
	time.Sleep(time.Second * 5)
	ticker.CancelTask(fmt.Sprint(10))
	time.Sleep(time.Hour)
}
