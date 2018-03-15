package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	//close(requests)
	go func() {
		close(requests)
	}()

	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	burstyLimter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimter <- t
		}
	}()

	burstyRequest := make(chan int, 5)
	for i:= 1; i<=5; i++ {
		burstyRequest <-i
	}

	close(burstyRequest)

	for req := range burstyRequest {
		<-burstyLimter
		fmt.Println("request", req, time.Now())
	}
}
