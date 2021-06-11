package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CalculateValue(values chan int) {
	rand.Seed(time.Now().Unix())
	value := rand.Intn(10)
	fmt.Println("Calculated Random Value: {}", value)
	values <- value
}

func main() {
	fmt.Println("Go Channel Tutorial")

	values := make(chan int)
	defer close(values)

	go CalculateValue(values)

	value := <-values
	fmt.Println(value)
}
