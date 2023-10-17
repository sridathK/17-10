package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	greet()
	fmt.Println("main")
}

func greet() {
	data := os.Args[1:]
	if len(data) != 3 {
		log.Println("please provide name , age , marks")
		os.Exit(1)
		//return // stops the exec of the current func
	}
	name := data[0]
	ageString := data[1]
	age, err := strconv.Atoi(ageString)
	marks, err := strconv.Atoi(data[2])
	if err != nil {
		// there is some kind of error
		log.Println(err)
		os.Exit(1)
		//return
	}
	fmt.Println(name, age, marks)

	//marksString := data[2]

}
