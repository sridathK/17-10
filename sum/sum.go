package sum

import "fmt"

var Sum int

func Add() {
	a, b := 10, 20
	Sum = a + b
	fmt.Println("this is add function")
	calcAll()
}
