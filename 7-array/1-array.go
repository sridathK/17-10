package main

import "fmt"

func main() {
	// var a [5]int
	// fmt.Println(a)
	// b := [5]int{10, 20, 30, 40, 50}
	// fmt.Println(b)
	// c := [...]int{10, 20, 50, 50, 60, 60, 60}
	// fmt.Println(len(c))

	// var i []int
	// fmt.Printf("%#v", i)

	// var i []int
	// i[1] = 100 // update
	// fmt.Println(i)

	var i []int
	//i[1] = 100 // update
	fmt.Println(i)
	if i == nil {
		fmt.Println("slice is nil", i)
	}
	i = append(i, 20, 30)
	fmt.Println(i)
	b := []int{10, 230, 4050, 60}
	fmt.Println(b)

}
