package main

import (
	"fmt"
	"sync"
)

func main() {
	var syncMap sync.Map = sync.Map{}
	var str string = "hello"

	for _, v := range str {
		var temp, ok = syncMap.Load(v)

		if !ok {
			syncMap.Store(v, 1)
		} else {
			syncMap.Store(v, temp.(int)+1)
		}
	}

	syncMap.Range(func(key, value interface{}) bool {
		fmt.Printf("key: %c, value: %v \n", key, value)

		return true
	})
}
