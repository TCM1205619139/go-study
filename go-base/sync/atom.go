package main

import "sync/atomic"

func main() {
	var newValue int32 = 1
	var dist int32 = 2
	var oldValue int32 = atomic.SwapInt32(&newValue, dist)

	println(oldValue)
}
