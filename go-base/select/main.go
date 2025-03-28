package main

import "fmt"

func main() {
	size := 10
	chan1 := make(chan int, size)
	chan2 := make(chan int, size)

	go func() {
		for i := 0; i < size; i++ {
			chan1 <- i
			chan2 <- i
		}
	}()

	for i := 0; i < size; i++ {
		select {
		case v1 := <-chan1:
			fmt.Printf("receive %d from channel 1\n", v1)
		case v2 := <-chan2:
			fmt.Printf("receive %d from channel 2\n", v2)
		}
	}
}
