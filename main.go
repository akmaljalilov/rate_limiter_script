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

func (sw *SlidingWindow) Allow() bool {

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
	c := float32(count)
	c /= float32(sw.timeUnit)
	c += float32(sw.currCount)

	if c > float32(sw.capacity) {
		fmt.Println("Dropped")
		return false
	}

	sw.currCount++
	fmt.Println("Passed")
	return true
}

func main() {
	//sw := NewSlidingWindow(2, 3)
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
	//time.Sleep(2 * time.Second)
	//for i := 0; i < 100; i++ {
	//	time.Sleep(100 * time.Millisecond)
	//	sw.Allow()
	//}
	sw := algorithms.New(3, 5)
	for i := 1; i < 100; i++ {
		sw.Allow()
		time.Sleep(time.Second)
	}
	//println(math.RoundToEven(6.5))
}
