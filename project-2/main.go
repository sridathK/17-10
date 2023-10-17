package main

import (
	"fmt"
	"log"
	"project-2/db"
)

func main() {
	a, b := db.NewConnection("raju")
	fmt.Println(a, b)
	u := db.Connection{Db: "ramu"}
	fmt.Println(u)
	log.Logger

}
