package main

import (
	"fmt"
)

func main() {

	ch := make(chan int) //unbuffered channel

	go func() {

		ch <- 20 // send will block until no receiver is ready
	}()

	x := <-ch //it is a blocking call until we don't recv the value
	fmt.Println(x)

	fmt.Println("end of main")

}
