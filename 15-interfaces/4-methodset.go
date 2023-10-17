package main

import "fmt"

type speaker interface {
	Speak() string
	Speak2() string
}

type Person struct {
	name string
}

func (p Person) Speak() string {
	return fmt.Sprintf("Hi, my name is %s", p.name)
}
func (p Person) Speak2() string {
	return fmt.Sprintf("Hi, my name is 2 %s", p.name)
}

func main() {
	p := Person{}
	p.Speak()
	var s1 speaker = Person{"John"} // This is valid
	fmt.Println(s1.Speak())
	fmt.Println(s1.Speak2())
}

//The method set for a value, only includes methods implemented with a value receiver.
//The method set for a pointer, includes methods implemented with both pointer and value receivers.
//The rules of method sets apply to interface types.
