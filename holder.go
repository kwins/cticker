package cticker

import (
	"sync"
)

// TaskHolder 持有所有的task
type TaskHolder struct {
	locker sync.RWMutex
	bulk   map[string]*Task
}

// NewTaskHolder NewTaskHolder
func NewTaskHolder() *TaskHolder {
	this := new(TaskHolder)
	this.bulk = make(map[string]*Task)
	return this
}

// Get Get
func (holder *TaskHolder) Get(sequenceid string) *Task {
	holder.locker.Lock()
	defer holder.locker.Unlock()
	if v, ok := holder.bulk[sequenceid]; ok {
		delete(holder.bulk, sequenceid)
		return v
	}
	return nil
}

// Add Add
func (holder *TaskHolder) Add(sequenceid string, task *Task) {
	holder.locker.Lock()
	holder.bulk[sequenceid] = task
	holder.locker.Unlock()
}

// Delete 删除
func (holder *TaskHolder) Delete(sequenceid string) {
	holder.locker.Lock()
	delete(holder.bulk, sequenceid)
	holder.locker.Unlock()
}
