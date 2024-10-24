package main

import (
	"fmt"
)

func main() {
	s := make([]int, 0, 2)

	doSomething(s)
	fmt.Println(s)
}

func doSomething(a []int) {
	a = append(a, 1)
	fmt.Println(a)
}
