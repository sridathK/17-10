package main

import "fmt"

//var dictionary = make(map[int][]string)

func main() {
	dictionary := make(map[int][]string)
	for k, v := range dictionary {
		fmt.Println(k, v)
	}
	dictionary[1] = []string{"cricket", "chess"}
	dictionary[2] = []string{"pubg", "carrom"}
	delete(dictionary, 1)
	//fmt.Println(dictionary["up"])
	for k, v := range dictionary {
		fmt.Println(k, v)
	}
	// fmt.Println(len(dictionary))

	//	search(3)
}

// func search(k1 int) {
// 	h, ok := dictionary[k1]
// 	if !ok {
// 		fmt.Println("not present")
// 		return
// 	}
// 	fmt.Println(h)
// }
