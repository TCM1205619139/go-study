package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
	member int
}

var instance *Singleton
var wg sync.WaitGroup
var once sync.Once

func getInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
		instance.member = 100
	})

	return instance
}

func main() {
	var size int = 10
	wg.Add(size)

	for i := 0; i < size; i++ {
		go func() {
			defer wg.Done()
			instance = getInstance()
		}()
	}
	wg.Wait()
	fmt.Println("All workers are done")
}
