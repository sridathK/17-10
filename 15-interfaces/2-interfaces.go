package main

import "fmt"

type function interface {
	info()
}

type zap struct {
	path string
}

type logrus struct {
	path string
}

func (z zap) info() {
	fmt.Println("this is terminal")
}
func (l logrus) info() {
	fmt.Println("this is file")
}

func doLog(f function) {
	f.info()

}

func main() {
	z := zap{path: "terminal"}
	l := logrus{path: "file"}
	doLog(z)
	doLog(l)
}
