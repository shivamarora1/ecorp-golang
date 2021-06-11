package main

import (
	"fmt"
	"time"
)

func runner1() {
	time.Sleep(time.Second * 3)
	fmt.Println("\nI am first runner")
}

func runner2() {
	fmt.Println("\nI am second runner")
}

func execute() {
	go runner1()
	go runner2()

}

func main() {

	// Launching both the runners
	execute()
	time.Sleep(time.Second * 1)

	fmt.Println("BYE BYE")
}
