package main

import "fmt"

type User struct {
	id       int
	name     string
	password string
}

var usermap = make(map[int]User)

func (u User) AddUser(id int) {
	usermap[id] = u
}

func main() {
	u1 := User{id: 1, name: "raju", password: "raju"}
	u2 := User{id: 2, name: "ramu", password: "ramu"}
	u1.AddUser(u1.id)
	u2.AddUser(u2.id)

	fmt.Println(usermap)

}
