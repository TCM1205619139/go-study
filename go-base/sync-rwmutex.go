package main

import "sync"

type Counter struct {
	value   int
	rwmutex sync.RWMutex
}

func (counter *Counter) getCounter() int {
	counter.rwmutex.RLock()
	defer counter.rwmutex.RUnlock()

	return counter.value
}

func (counter *Counter) increment(value int) {
	counter.rwmutex.Lock()
	defer counter.rwmutex.Unlock()

	counter.value += value
}

var wg sync.WaitGroup

func main() {
	size := 100
	wg.Add(size)
	counter := Counter{value: 0}

	for i := 0; i < size; i++ {
		go func(v int) {
			defer wg.Done()

			counter.increment(v)
		}(i)
	}

	wg.Wait()
	println(counter.getCounter())
}
