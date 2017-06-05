package cticker

import (
	"sync"
)

// taskHolder 持有所有的task
type taskHolder struct {
	locker sync.RWMutex
	bulk   map[string]*Task
}

// newTaskHolder newTaskHolder
func newTaskHolder() *taskHolder {
	this := new(taskHolder)
	this.bulk = make(map[string]*Task)
	return this
}

// Get Get
func (holder *taskHolder) get(sequenceid string) *Task {
	holder.locker.Lock()
	defer holder.locker.Unlock()
	if v, ok := holder.bulk[sequenceid]; ok {
		delete(holder.bulk, sequenceid)
		return v
	}
	return nil
}

// Add Add
func (holder *taskHolder) add(sequenceid string, task *Task) {
	holder.locker.Lock()
	holder.bulk[sequenceid] = task
	holder.locker.Unlock()
}

// Delete 删除
func (holder *taskHolder) delete(sequenceid string) {
	holder.locker.Lock()
	delete(holder.bulk, sequenceid)
	holder.locker.Unlock()
}

func (holder *taskHolder) cancel(sequenceid string) {
	holder.get(sequenceid).cancel = true
}
