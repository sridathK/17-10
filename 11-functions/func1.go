package main

import (
	"fmt"
)

func main() {
	add := func(a, b int) int { return a + b }
	fmt.Println(add(1, 2))
}
func Operate(op func(x, y int) int, i int) {
	fmt.Println(op(1, 4))
	fmt.Println(i)
}
