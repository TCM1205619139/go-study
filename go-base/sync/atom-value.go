package main

import (
	"sync/atomic"
	"time"
)

func loadConfig() map[string]string {
	return make(map[string]string)
}

func requests() chan int {
	return make(chan int)
}

func main() {
	var config atomic.Value

	config.Store(loadConfig())

	go func() {
		time.Sleep(10 * time.Second)
		config.Store(loadConfig())
	}()

	for i := 0; i < 10; i++ {
		go func() {
			for r := range requests() {
				c := config.Load()
				_, _ = r, c.(map[string]string)
			}
		}()
	}
}
