package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%T", os.Args)
	// fmt.Println(os.Args[0])
	// fmt.Println(os.Args[1])
	// fmt.Println(os.Args[2])
	a := os.Args[1:]
	fmt.Println(a)
}
