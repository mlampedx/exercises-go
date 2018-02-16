package main

import (
	"fmt"
	"time"

	"github.com/mlampedx/exercises-go/limiter/ticker"
)

func main() {
	reqs := make(chan int, 5)
	defer close(reqs)
	for i := 1; i <= 5; i++ {
		reqs <- i
	}

	limiter := time.Tick(200 * time.Millisecond)
	for req := range reqs {
		<-limiter
		fmt.Println("request", req, "at", time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i <= 3; i++ {
		burstyLimiter <- time.Now()
	}

	go ticker.Create(burstyLimiter)

	burstyReqs := make(chan int, 5)
	defer close(burstyReqs)
	for i := 1; i <= 5; i++ {
		burstyReqs <- i
	}

	for req := range burstyReqs {
		<-burstyLimiter
		fmt.Println("request", req, "at", time.Now())
	}
}
