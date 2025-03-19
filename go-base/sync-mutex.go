package main

import (
	"fmt"
	"sync"
)

var num int = 0
var wg sync.WaitGroup
var mutex sync.Mutex

func add(value int) {
	mutex.Lock()
	defer mutex.Unlock()
	num += value
}

func main() {
	size := 100
	wg.Add(size)

	for i := 0; i < size; i++ {
		i := i
		go func() {
			defer wg.Done()
			add(i)
		}()
	}

	wg.Wait()
	fmt.Printf("sum of 1 to %d is: %d\n", size, num)
}
