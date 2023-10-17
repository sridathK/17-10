package main

import (
	"fmt"
	"strconv"
)

// func details(a string, b int) {
// 	fmt.Printf(a+"age is %p", b)
// }

// func Name(a, b string) string {
// 	return a + b
// }

func square(a, b string) (int, error) {
	c, err := strconv.Atoi(a)
	if err != nil {
		return 0, err
	}
	d, err := strconv.Atoi(b)
	if err != nil {
		return 0, err
	}
	return c + d, nil

}
func main() {
	// details("ramu", 20)
	// fmt.Println("ramu", "is king")
	a, err := square("2e", "3")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(a)

}
