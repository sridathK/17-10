package main

import "fmt"

type User struct {
	name string
	age  int
}

func main() {
	u := []User{{name: "raju", age: 2}, {name: "ramu", age: 3}}

	fmt.Println(u)

}
