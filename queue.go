package cticker

import (
	"sync"
	"time"

	log "github.com/cihub/seelog"
)

const defaultSlotNum = 512
const defaultDuration = time.Second

// Queue 环形队列
type Queue struct {
	slotNum   int          // 环形队列的Slot数量
	s         *Slots       // 环形队列的槽
	taskHoler *TaskHolder  // 环形队列所有的定时任务
	locker    sync.RWMutex //
}

// NewQueue 新建一个num个slot的环形队列
// 环形队列序号从 1 开始
func NewQueue(num int, duration time.Duration) *Queue {
	q := new(Queue)
	q.s = NewSlots(num, duration)
	q.taskHoler = NewTaskHolder()
	q.slotNum = cap(q.s.slots)
	q.s.Loop()
	return q
}

// GetTask 获取一个定时任务
func (q *Queue) GetTask(sequenceid string) *Task {
	return q.taskHoler.Get(sequenceid)
}

// CancelTask 取消尚未执行的定时任务
func (q *Queue) CancelTask(sequenceid string) {
	if t := q.taskHoler.Get(sequenceid); t != nil {
		t.Cancel()
	}
}

// AddTimerTask 增加定时任务
func (q *Queue) AddTimerTask(seconds int, sequenceid string, task *Task) error {

	count := q.s.current + seconds
	task.cycleNum = count / q.slotNum
	index := count % q.slotNum
	log.Trace("current:", q.s.current, " count:", q.s.count, " count:", count, " task.cycleNum:", task.cycleNum, " index:", index)
	q.taskHoler.Add(sequenceid, task)

	return q.s.AddByIndex(index, task)
}
