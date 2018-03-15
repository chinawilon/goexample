package main

import (
	"fmt"
	//"time"
)

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all job")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("send job", j)
	}

	close(jobs)
	fmt.Println("send all jobs")

	//time.Sleep(2 * time.Second)
	//fmt.Scanln()
	<-done
}
