package main

import (
	"PROJECT-3/stores"
	"PROJECT-3/stores/Mysql"
	//"PROJECT-3/stores/postgres"
)

func main() {
	u := stores.User{
		Name:  "ajay",
		Email: "ajay@email.com",
	}

	m := Mysql.New("mysql")
	// stores.StorerInterface = m
	// stores.StorerInterface.Create(u)

	// p := postgres.New("postgre")
	// stores.StorerInterface = p
	// stores.StorerInterface.Create(u)
	stores.NewService(m).Storer.Create(u)

}
