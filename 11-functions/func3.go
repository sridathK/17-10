package main

import "fmt"

type operation func(x, y int) int

func main() {
	s := func(a, b int) int {
		return a - b
	}
	operate(add, 20, 30)
	operate(s, 40, 30)

}
func operate(op operation, a, b int) {
	fmt.Println(op(10, 20))
}
