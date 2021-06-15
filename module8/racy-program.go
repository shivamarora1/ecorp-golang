package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wgIns sync.WaitGroup

var (
	counter int = 0
	m       sync.Mutex
)

func incr() {

	for j := 0; j < 10; j++ {

		m.Lock()
		counter = counter + 1
		m.Unlock()
	}
	wgIns.Done()
}

func main() {

	wgIns.Add(100)
	for i := 0; i < 100; i++ {
		go incr()

		//Iteration: 1; Counter:10
		//Iteration: 2; Counter:20
		//	.....
		//Iteration: 10; Counter:100

	}

	fmt.Println("The number of goroutines before wait = ", runtime.NumGoroutine())
	wgIns.Wait()

	fmt.Println("Counter value = ", counter)
	fmt.Println("The number of goroutines after wait = ", runtime.NumGoroutine())
}
