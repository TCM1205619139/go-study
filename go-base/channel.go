package main

import (
	"fmt"
	"time"
)

type Cat struct {
	name string
	age  int
}

func fetchChannel(ch chan Cat) {
	value := <-ch
	fmt.Printf("type: %T, value: %v\n", value, value)
}

func main() {
	ch := make(chan Cat, 100)
	go fetchChannel(ch)
	ch <- Cat{"yingduan", 1}
	time.Sleep(2 * time.Second)
	fmt.Println("end")
}
