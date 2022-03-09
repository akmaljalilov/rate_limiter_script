package algorithms

import (
	"fmt"
	"time"
)

type FixedWindow struct {
	windows          map[int]int
	maxRequestPerSec int
}

func fixedWindow(maxRequestPerSec int) *FixedWindow {
	return &FixedWindow{
		windows:          make(map[int]int),
		maxRequestPerSec: maxRequestPerSec,
	}
}
func (b *FixedWindow) Allow() bool {
	curSec := time.Now().Unix()
	b.windows[int(curSec)]++
	fmt.Println(curSec)
	if b.windows[int(curSec)] <= b.maxRequestPerSec {
		fmt.Println("Passed")
		return true
	}
	fmt.Println("Dropped")
	return false
}
