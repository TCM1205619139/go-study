package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	var size int = 10
	wg.Add(size)

	for i := 0; i < size; i++ {
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	wg.Wait()
	fmt.Println("end")
}
