package main

import "fmt"

func main() {
	dictionary := make(map[string]string)

	dictionary["below"] = "down"
	dictionary["up"] = "above"
	fmt.Println(dictionary["up"])
	for k, v := range dictionary {
		fmt.Println(k, v)
	}
	fmt.Println(len(dictionary))
}
