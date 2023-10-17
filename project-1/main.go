package main

import (
	"PROJECT-1/models"
	"fmt"
)

func main() {

	u := models.User{Name: "raj", Email: "12@gmail",
		Permissions: map[string]bool{"admin": true}}

	fmt.Println(u)
	for k, v := range u.Permissions {
		fmt.Println(k, v)
	}
}
