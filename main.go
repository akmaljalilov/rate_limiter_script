package main

import (
	"fmt"
	"time"
	"velox/rate_limiter/algorithms"
)

type SlidingWindow struct {
	timeUnit    int64 // the window duration
	capacity    int64 // request capacity
	currentTime int64 // current window time
	prevCount   int64 // previous count of requests
	currCount   int64 // how many requests are processed in current window
}

func NewSlidingWindow(capacity, duration int64) *SlidingWindow {
	return &SlidingWindow{
		capacity:    capacity,
		timeUnit:    duration,
		currentTime: time.Now().Unix(),
		prevCount:   capacity,
		currCount:   0,
	}
}

func (sw *SlidingWindow) Handle(packet interface{}) bool {

	now := time.Now().Unix()
	count := now - sw.currentTime

	// check if window is valid
	if count > sw.timeUnit {
		sw.currentTime = now
		sw.prevCount = sw.capacity
		sw.currCount = 0
	}

	count = sw.timeUnit - count
	if count < 0 {
		count = 0
	}
	count *= sw.prevCount
	count /= sw.timeUnit
	count += sw.currCount

	if count > sw.capacity {
		fmt.Println("Dropped")
		return false
	}

	sw.currCount++
	fmt.Println("Passed")
	return true
}

func main() {
	sw := algorithms.New(2, 3)
	//sw := NewSlidingWindow(2, 1)
	//time.Sleep(2 * time.Second)
	/*i := 0
	for {
		i++
		if i == 1000 {
			break
		}
		time.Sleep(1 * time.Second)
		if sl.Allow() {
			println(i)
		}
	}*/
	for i := 0; i < 100; i++ {
		time.Sleep(1 * time.Second)
		if sw.Allow() {
			println(i)
		}
	}
}
