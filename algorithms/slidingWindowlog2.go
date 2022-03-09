package algorithms

import (
	"fmt"
	"time"
)

type SlidingWindow struct {
	capacity  int64
	timeUnit  int64
	curTime   int64
	curCount  int64
	prevCount int64
}

func New(cap int64, timeUnit int64) *SlidingWindow {
	return &SlidingWindow{
		capacity:  cap,
		timeUnit:  timeUnit,
		curTime:   time.Now().Unix(),
		curCount:  0,
		prevCount: 0,
	}
}

func (s *SlidingWindow) Allow() bool {
	cur := time.Now().Unix()
	r:=cur-s.curTime
	if r > s.timeUnit {
		s.curTime = cur
		s.prevCount = s.curCount
		s.curCount = 0
	}
	ec := (s.prevCount * (s.timeUnit - r) / s.timeUnit) + s.curCount
	if ec >= s.capacity {
		fmt.Println("Dropped")
		return false
	}
	s.curCount++
	fmt.Println("Passed")
	return true
}
