package algorithms

import (
	"fmt"
	"math"
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
		prevCount: cap,
	}
}

func (s *SlidingWindow) Allow() bool {
	cur := time.Now().Unix()
	r := cur - s.curTime
	if r > s.timeUnit {
		s.curTime = cur
		s.prevCount = s.curCount - 1
		s.curCount = 0
	}
	nisbat := float64(s.timeUnit-(cur-s.curTime)) / float64(s.timeUnit)
	pr := int64(math.Trunc(float64(s.prevCount) * nisbat))
	ec := pr + s.curCount
	if ec > s.capacity {
		fmt.Println("Dropped", ec, pr, (float64(s.prevCount) * nisbat))
		return false
	}
	fmt.Println("Passed", ec, pr, (float64(s.prevCount) * nisbat))
	s.curCount++
	return true
}
