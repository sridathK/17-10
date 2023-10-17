package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

var ErrFileNotFound = errors.New("not in the root directory")

func main() {
	_, err := openFile("test.txt")
	if err != nil {
		log.Println(err)
		if errors.Is(err, ErrFileNotFound) {
			fmt.Println("present or found in chain")
			return
		}
		return
	}
}

func openFile(fileName string) (*os.File, error) {
	f, err := os.Open(fileName)

	if err != nil {
		return nil, fmt.Errorf("%w %w", err, ErrFileNotFound)
	}
	return f, nil
}
