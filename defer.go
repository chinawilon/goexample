package main

import (
	"fmt"
	"os"
)

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprint(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}

func main() {
	f := createFile("tem.sql")
	defer closeFile(f)
	writeFile(f)
}
