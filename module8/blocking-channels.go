package main

import (
	"time"
)

//"fmt"

func run(c chan int) {
	c <- 34
}

func main() {
	c := make(chan int)
	go run(c)

	make(chan int, 3)

	time.Sleep(45 * time.Second)
	//fmt.Printf("%v", <-c)
}
