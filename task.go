package cticker

import (
	"container/list"
	"sync"
)

// Task Slot中定时任务个体
type Task struct {
	seqid    string       //
	cycleNum int          // 第几圈执行此定时任务
	handler  func() error // 执行任务的函数,参数个自定义
	cancel   bool         // true-表示取消，false-表示执行
}

// SetTaskHandler 设置任务函数
func (task *Task) SetTaskHandler(h func() error) {
	if nil == h {
		task.handler = func() error {
			return nil
		}
	}
	task.handler = h
}

// Cancel 取消任务
func (task *Task) Cancel() {
	task.cancel = true
}

// Tasks Slot中任务集合
type Tasks struct {
	locker sync.RWMutex
	tasks  *list.List // Slot中任务集合
}

// Remove 移除链表元素
func (t *Tasks) Remove(e *list.Element) {
	t.locker.Lock()
	t.tasks.Remove(e)
	t.locker.Unlock()
}

// PushBack 加入到链表最后
func (t *Tasks) PushBack(task *Task) {
	t.locker.Lock()
	t.tasks.PushBack(task)
	t.locker.Unlock()
}
