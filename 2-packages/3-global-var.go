package main

import (
	"fmt"
	"go-training/sum"
	"os"
)

var s = "jfdkd" +
	"kcmskcldmcld"

func main() {
	sum.Add()
	fmt.Println(s)
	sum.Sum = 33
	fmt.Println(sum.Sum)
	os.Args
}
