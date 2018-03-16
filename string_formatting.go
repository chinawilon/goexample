package main

import "fmt"

type Pointer struct {
	X, Y int
}

func main() {
	p := Pointer{1, 2}
	fmt.Println(p)
	fmt.Printf("%v\n", p)
	fmt.Printf("%V\n", p)
	fmt.Printf("%+v\n", p)
	fmt.Printf("%#v\n", p)

	fmt.Printf("%t\n", p)
	fmt.Printf("%T\n", p)

	fmt.Printf("%s\n", "\"string\"")
	fmt.Printf("%q\n", "\"string\"")
	fmt.Printf("%x\n", "hex this")
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)
}
