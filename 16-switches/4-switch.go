package main

import (
	"fmt"
)

type Man struct {
	name string
}

type runner interface {
	run() int
}

type walker interface {
	walk()
}
type sprinter interface {
	runner
	walker
}

func (m Man) run() int {
	return 1
}
func (m Man) walk() {

}

func main() {
	m := Man{name: "ra"}
	//var r runner = m
	//var w walker = m
	var wr sprinter = m
	fmt.Println(wr.run())

}
