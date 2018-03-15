package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}
func (r *rect) permi() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{10, 20}

	fmt.Println("area:", r.area())
	fmt.Println("permi:", r.permi())

	rp := &r

	fmt.Println("area:", rp.area())
	fmt.Println("permi", rp.permi())
}
