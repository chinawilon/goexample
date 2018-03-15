package main

import (
	"fmt"
)

func f(from string) {
	for i := 0; i < 100; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	f("deirect")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	fmt.Println("done")

}
