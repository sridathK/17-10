package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		//time.Sleep(4)
		ch1 <- 10
	}()

	go func() {
		ch2 <- 20
	}()

	for i := 0; i < 2; i++ {
		select {
		case x := <-ch1:
			fmt.Println(x) //select staemnt

		case y := <-ch2:
			fmt.Println(y)
		}
	}

	fmt.Println("end of main")

}
