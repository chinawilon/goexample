package main

import "os"

func main() {

	//panic("a problem")

	_, err := os.Create("file.sql")
	if err != nil {
		panic(err)
	}
}
