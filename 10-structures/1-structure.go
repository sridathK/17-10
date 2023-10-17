package main

import (
	"fmt"
)

type money int //new type is money and underlying type is int

func main() {
	// var Rupe money = 10
	// fmt.Printf("%T", Rupe)
	// //fmt.Println(time.Duration)
	// i := 20
	// Rupe = money(i)
	//                                                        //types.
	// var a int64
	// fmt.Println()
	// fmt.Println(a)
	u := User{Name: "raj", Age: 20, Marks: []int{22, 34}}
	fmt.Printf("%+v", u)
	fmt.Println()
	fmt.Printf("%#v", u)
	//u1 := User{Name: "raj", Age: 20,Marks: []int{22,34}}
	// if u == u1 {
	// 	fmt.Println("yes")
	// }

}

type User struct {
	Name  string //structure
	Age   int
	Marks []int
}
