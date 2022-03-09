package algorithms

import (
	"fmt"
	"time"
)

type LeakyBucket struct {
	nextAllowedTime        int64
	requestIntervalSeconds int64
}

func leakyBucket(sec time.Duration, maxRequestPerSec int64) *LeakyBucket {
	return &LeakyBucket{
		nextAllowedTime:        time.Now().UnixMilli(),
		requestIntervalSeconds: int64(sec/time.Millisecond) / maxRequestPerSec,
	}
}

func (b *LeakyBucket) Allow() bool {
	curTime := time.Now().UnixMilli()
	if curTime >= b.nextAllowedTime {
		b.nextAllowedTime += b.requestIntervalSeconds
		fmt.Println("Passed")
		return true
	}
	fmt.Println("Dropped")
	return false
}
