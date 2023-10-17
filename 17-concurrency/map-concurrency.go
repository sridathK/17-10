package main

import (
	"fmt"
	"sync"
)

var myMap = make(map[int]int)
var wg = &sync.WaitGroup{}
var m = &sync.Mutex{}

func main() {
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			m.Lock()
			defer m.Unlock()
			myMap[n] = n * n

		}(i)
	}
	wg.Wait()
	fmt.Println(myMap)
}
