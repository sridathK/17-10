package main

import (
	"fmt"
	"time"
)

func main() {
	go hello()                     // spawning a goroutine //go before method, ready to run,but not running.
	fmt.Println("end of the main") //
	time.Sleep(2 * time.Second)
}

func hello() {
	fmt.Println("hello from the hello func")
}
