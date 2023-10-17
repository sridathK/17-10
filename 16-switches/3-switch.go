package main

import "io"

type student struct {
}

func (s student) Write(p []byte) (n int, err error) {
	//TODO implement me
	panic("implement me")
}

func (s student) Read(p []byte) (n int, err error) {
	//TODO implement me
	panic("implement me")
}

func (s student) Walk() {
}

func (s student) Run() {

}

func main() {
	var s student
	var r io.Reader = s
	var w io.Writer = s
	var rw io.ReadWriter = s
	io.ReadWriter
	_, _, _ = r, w, rw
}
