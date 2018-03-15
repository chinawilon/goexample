package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "b", "a", "d"}
	sort.Strings(strs)
	fmt.Println(strs)
	ints := []int{1, 2, 5, 6, 3, 9}
	sort.Ints(ints)
	fmt.Println(ints)
	s := sort.IntsAreSorted(ints)
	fmt.Println(s)
}
