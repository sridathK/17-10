package postgres

import (
	"PROJECT-3/stores"
	"fmt"
)

type Conn struct {
	db string
}

func New(db string) Conn {
	return Conn{db: db}
}

func (c Conn) Create(usr stores.User) error {
	fmt.Println("user create")
	return nil
}
func (c Conn) Update(usr stores.User) error {
	fmt.Println("user update")
	return nil
}

func (c Conn) Delete(usr stores.User) error {
	fmt.Println("user delete")
	return nil
}
