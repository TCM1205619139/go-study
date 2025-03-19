package main

import (
	"fmt"
	"sync"
	"time"
)

func print() {
	fmt.Println("Hello, World!")
}

func main() {
	size := 10
	var wg sync.WaitGroup
	var once sync.Once
	wg.Add(size)

	for i := 0; i < size; i++ {
		go func(id int) {
			defer wg.Done()
			time.Sleep(time.Second)
			once.Do(print)
		}(i)
	}

	wg.Wait()
	fmt.Println("end")
}
