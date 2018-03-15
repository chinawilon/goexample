package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nexInt := intSeq()
	fmt.Println(nexInt())
	fmt.Println(nexInt())
	fmt.Println(nexInt())

	newInt := intSeq()
	fmt.Println(newInt())
}
