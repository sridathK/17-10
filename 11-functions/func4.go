package main

import "fmt"

func main() {
	request(get, "get")

}
func get(endpoint string) {
	fmt.Println("this is get method")
}

func post(endpoint string) {
	fmt.Println("this is post method")
}

func put(endpoint string) {
	fmt.Println("this is put method")
}
func request(f func(x string), endpoint string) {
	f(endpoint)
}
