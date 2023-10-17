package main

import (
	"fmt"
)

type users struct {
	name string
}

func main() {
	checkType(users{name: ""})
}

func checkType(i any) {
	switch i.(type) {
	case int:
		fmt.Println("it is an int value")

	case users:
		fmt.Println("this is user")
	default:
		fmt.Println("unknown  type")
	}

}
