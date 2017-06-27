package cticker

import (
	"sync"
	"time"
)

const defaultSlotNum = 512
const defaultDuration = time.Second

// Queue 环形队列
type Queue struct {
	slotNum   int          // 环形队列的Slot数量
	s         *slots       // 环形队列的槽
	taskHoler *taskHolder  // 环形队列所有的定时任务
	locker    sync.RWMutex //
}

// NewQueue 新建一个num个slot的环形队列
// 环形队列序号从 1 开始
func NewQueue(num int, duration time.Duration) *Queue {
	q := new(Queue)
	q.s = newSlots(num, duration)
	q.taskHoler = newTaskHolder()
	q.slotNum = cap(q.s.slots)
	q.s.loop()
	return q
}

// GetTask get un exec task
func (q *Queue) GetTask(sequenceid string) *Task {
	return q.taskHoler.get(sequenceid)
}

// CancelTask 取消尚未执行的定时任务
func (q *Queue) CancelTask(sequenceid string) {
	q.taskHoler.cancel(sequenceid)
}

// AddTimerTask 增加定时任务
func (q *Queue) AddTimerTask(seconds int, sequenceid string, task *Task) error {
	task.seqid = sequenceid
	count := q.s.current + seconds
	task.cycleNum = count / q.slotNum
	index := count % q.slotNum
	q.taskHoler.add(sequenceid, task)

	return q.s.addByIndex(index, task)
}
