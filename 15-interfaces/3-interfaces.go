package main

import "fmt"

type user1 struct {
	name string
}

func main() {
	var i interface{} = "str"
	//var a any // any is an alias to interface{}
	i = 10
	i = true
	i = "hello"
	var u user1
	i = u
	str, ok := i.(int) // type assertion
	if !ok {
		fmt.Println("user type is not there")
		return
	}
	fmt.Println(str)
}
