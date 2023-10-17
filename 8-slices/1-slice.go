package main

import "fmt"

func main() {

	// i := []int{10, 20, 24}
	// j := i[:2]
	// fmt.Println(i)
	// fmt.Println(j)

	// a := []int{10, 20, 30, 40, 50, 60}
	// b := a[1:4]
	// fmt.Println(b)                                    //for referencing.
	// b[0] = 777
	// fmt.Println(a[0])
	// fmt.Println(a)

	// a := []int{10, 20, 30, 40, 50, 60}       //here length and capacity are equal.
	// b := a[2:5]
	// b = append(b, 999)                              // append
	// fmt.Println(a)
	// fmt.Println(b)

	// a := []int{10, 20, 30, 40, 50}
	// var b []int
	// i := a[1:5]
	// fmt.Println("capacity of i ", cap(i))
	// fmt.Println("length of i", len(i))
	// fmt.Printf("%p", a)
	// a = append(a, 60)
	// fmt.Println()
	// fmt.Printf("%p", a)
	// fmt.Println()
	// // fmt.Println(i)
	// // fmt.Println(a)
	// // fmt.Println("capacity of appended i ", cap(i))
	// // fmt.Println("length of appended i", len(i))
	// fmt.Printf("%p", b)
	// fmt.Println(b)
	// if len(b) == 0 {
	// 	fmt.Println(222)
	// }

	// a := make([]int, 2, 2) // length,capacity.
	// fmt.Println(len(a))
	// fmt.Println(cap(a))
	// fmt.Println(a)                              //preallocating a backing array.
	// fmt.Printf("%p", a)
	// a = append(a, 100)
	// fmt.Println()
	// fmt.Printf("%p", a)

	// a := []int{1, 2, 3, 4}
	//b := make([]int, 0, 100)
	// copy(b, a)
	// fmt.Println(a)                                           //copying
	// fmt.Println(b)                                           //make and copy  works together.
	// //b[0] = 100
	// b = append(b, 20)
	// fmt.Println(a)
	// fmt.Println(b)

	a := []int{10, 20, 30, 40, 50}
	var b []int
	//b := make([]int, len(a[1:3]))
	fmt.Printf("%p", b)
	fmt.Println(b)

	//copy(b, a[1:3]) //b=a is used for copying referencing.
	fmt.Println(b)

}
