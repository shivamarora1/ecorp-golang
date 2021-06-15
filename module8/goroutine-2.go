package main

import (
	"fmt"
	"sync"
	//"time"
)

var WG *sync.WaitGroup

// notice we've not changed anything in this function
// when compared to our previous sequential program
func compute(thread int, value int) {
	defer WG.Done()
	fmt.Println("Running Goroutine ", thread)

}

func main() {
	fmt.Println("Goroutine Tutorial")

	WG = new(sync.WaitGroup)
	WG.Add(8)
	// notice how we've added the 'go' keyword
	// in front of both our compute function calls
	go compute(1, 10)
	go compute(2, 10)
	go compute(3, 10)
	go compute(4, 10)
	go compute(5, 10)
	go compute(6, 10)
	go compute(7, 10)
	go compute(8, 10)

	WG.Wait()
}
