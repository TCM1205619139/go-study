package main

import (
	"fmt"
	"sync"
)

var syncMap sync.Map = sync.Map{}
var wg sync.WaitGroup

func changeMap(key int) {
	defer wg.Done()
	syncMap.Store(key, 1)
}

func main() {
	size := 10
	wg.Add(size)

	for i := 0; i < size; i++ {
		go func() {
			defer wg.Done()
			changeMap(i)
		}()
	}

	wg.Wait()
	syncMap.Range(func(key, value interface{}) bool {
		fmt.Printf("key: %v, value: %v \n", key, value)

		return true
	})
}
