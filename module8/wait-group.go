package main

import (
	"fmt"
	"sync"
	"time"
)

func runner1(wg *sync.WaitGroup) {
	fmt.Print("\nI am first runner")
	time.Sleep(3 * time.Second)
	wg.Done()
}

func runner2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Print("\nI am second runner")
}

func execute() {

	wg := new(sync.WaitGroup)
	wg.Add(4)

	go runner1(wg)
	go runner2(wg)

	wg.Wait()
}

func main() {
	execute()
	fmt.Printf("\n%v\n", "BYE BYE")
}
