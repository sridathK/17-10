package main

import "fmt"

type User struct {
	id       int
	name     string
	password string
}

func (u User) updateUser(id int) {
	u.id = id
	fmt.Println(u)
}

func main() {
	u1 := User{id: 12, name: "raju"}
	u1.updateUser(13)
	fmt.Println(u1)

}
