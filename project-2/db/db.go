package db

//import "fmt"

type Connection struct {
	Db string
}

func NewConnection(db string) (Connection, bool) {
	if db == "" {                                         //factory function.
		return Connection{}, false
	}
	return Connection{Db: db}, true

}
