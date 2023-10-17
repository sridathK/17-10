package main

import (
	"fmt"
	"log"
)

type user struct {
	name string
}

func (u user) Write(p []byte) (n int, err error) {
	fmt.Println("this is implementation")
	return 0,nil
}

func main() {
	u := user{name: "ajay"}
	l := log.New(u, "sales:", log.Lshortfile)
	l.Println()

}
