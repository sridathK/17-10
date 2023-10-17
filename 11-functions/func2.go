package main

import "fmt"

func main() {
	// operate(add)
	fmt.Println((sub(3, 2)))

}
func operate(op func(x, y int) int) {
	fmt.Println(op(10, 20))
}

// func add(a, b int) int {
//     return a + b
// }

func sub(a, b int) int {
	return a - b
}
