package main

import (
	"log"
	"os"
)

func main() {
	l := log.New(os.Stdout, "sales:", log.Lshortfile)
	l.Println()
}
