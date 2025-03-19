package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	size := 10
	var wg sync.WaitGroup
	var mutex sync.Mutex

	wg.Add(size)
	var cond sync.Cond = *sync.NewCond(&mutex)

	for i := 0; i < size; i++ {
		go func() {
			defer wg.Done()
			cond.L.Lock()
			fmt.Printf("%d ready\n", i)
			cond.Wait()
			fmt.Printf("%d done\n", i)
			cond.L.Unlock()
		}()
	}

	time.Sleep(2 * time.Second)

	go func() {
		defer wg.Done()
		cond.Broadcast()
	}()
	wg.Wait()
}
