package main

import (
	"errors"
	"fmt"
)

var ErrStringEmpty = errors.New("string is empty")

func FetchData(s string) (string, error) {
	if s == "" {
		return "", ErrStringEmpty
	}
	return s, nil

}

func main() {
	s, err := FetchData("r")
	fmt.Println(s, err)
}
