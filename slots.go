package cticker

import (
	"container/list"
	"errors"
	"log"
	"time"
)

var (
	errIndexOutofRange = errors.New("index out of range")
	errTaskNil         = errors.New("task nil")
)

// Slots 循环队列槽
type slots struct {
	slots    []*Tasks      //
	count    int           // 槽总数
	current  int           // 当前Slot index
	duration time.Duration // 槽移动单位时间
}

// newSlots new slots
// 设置rdb的值，则支持持久化存储模式，rdb为nil则为内存模式
func newSlots(count int, duration time.Duration) *slots {
	s := new(slots)
	if count != 0 {
		s.count = count
		s.slots = make([]*Tasks, count+1)
	} else {
		s.count = defaultSlotNum
		s.slots = make([]*Tasks, defaultSlotNum+1)
	}
	for k := range s.slots {
		s.slots[k] = new(Tasks)
		s.slots[k].tasks = list.New()
	}
	s.current = 1
	if duration == 0 {
		s.duration = defaultDuration
	} else {
		s.duration = duration
	}
	return s
}

// Loop 槽循环
func (s *slots) loop() {
	t := time.NewTicker(s.duration)
	go func() {
		for {
			<-t.C
			s.next()
		}
	}()
}

// Next Next
func (s *slots) next() {
	// 检查当前Slot中有没有定时任务需要执行
	for e := s.slots[s.current].tasks.Front(); e != nil; e = e.Next() {
		v, ok := e.Value.(*Task)
		if !ok {
			s.slots[s.current].Remove(e)
			log.Println("remove:", s.current)
			continue
		}
		if 0 == v.cycleNum {
			if !v.cancel {
				go v.handler()
			}
			s.slots[s.current].Remove(e)
		} else {
			v.cycleNum--
		}
	}
	// 移动到下一个Slot
	s.current++
	if s.current == s.count {
		s.current = 1
	}
}

// AddByIndex 增加定时任务
func (s *slots) addByIndex(index int, task *Task) error {
	if index > s.count {
		return errIndexOutofRange
	}
	if nil == task {
		return errTaskNil
	}
	s.slots[index].PushBack(task)
	return nil
}
