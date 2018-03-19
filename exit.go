package main

import (
	"fmt"
	//"os"
	"syscall"
)

func main() {

	defer fmt.Println("exiting")

	//os.Exit(4)
	syscall.Exit(2)

}
