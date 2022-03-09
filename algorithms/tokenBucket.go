package algorithms

import (
	"fmt"
	"time"
)

type TokenBucket struct {
	sec              time.Duration
	maxRequestPerSec int
	tokens           int
}

func tokenBucket(sec time.Duration, maxRequest int) *TokenBucket {
	t := &TokenBucket{
		sec:              sec,
		maxRequestPerSec: maxRequest,
		tokens:           maxRequest,
	}
	go t.refillTokens()
	return t

}
func (b *TokenBucket) Allow() bool {
	if b.tokens != 0 {
		b.tokens--
		fmt.Println("Passed")
		return true
	}
	fmt.Println("Dropped")
	return false
}

func (b *TokenBucket) refillTokens() {
	for range time.Tick(time.Second * b.sec) {
		b.tokens = b.maxRequestPerSec
	}
}
