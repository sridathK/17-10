package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	data := os.Args[1:]
	if len(data) != 3 {
		fmt.Println("please provide name,age,marks")
		return
	}
	name := data[0]
	age, err := strconv.Atoi(data[1])
	marks := data[2]
}
