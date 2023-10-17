package main

import (
	"errors"
	"fmt"
	"os"
)

type QueryErrorr struct {
	dirName string
	Err     error
}

var q *QueryError

func (q *QueryErrorr) Error() string {
	return q.dirName + " " + q.Err.Error()
}

func main() {
	_, err := openFilee("test11.txt", "root")
	// if err != nil {
	// 	log.Println(err)

	// }
	if err != nil {
		if errors.As(err, &q) {
			fmt.Println("custom error found in the chain", q.dirName)
			return
		}
	}

}

func openFilee(fileName string, dirName string) (*os.File, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, &QueryErrorr{
			dirName: "wrong1",
			Err:     err,
		}
	}
	return f, nil
}
