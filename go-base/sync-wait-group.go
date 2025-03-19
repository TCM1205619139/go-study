package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	fmt.Println("Worker is working; Worker ID: %d", id)
	time.Sleep(2 * time.Second)
	wg.Done()
	fmt.Println("Worker is done %d", id)
}

func main() {
	var wg sync.WaitGroup
	size := 10

	wg.Add(size)

	for i := 0; i < size; i++ {
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("All workers are done %v", wg)
}
