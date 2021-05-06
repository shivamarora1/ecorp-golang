//For loop example
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// From the given trees we have to plant
	// different trees according to given number
	rand.Seed(time.Now().Unix())
	availableTrees := []string{"Banana", "Oak", "Papaya", "Orange", "Melon"}
	land := make(map[string]int)
	treesToPlant := 23

	for i := 0; i < treesToPlant; i++ {
		r := rand.Intn(len(availableTrees))
		tree := availableTrees[r]

		if _, ok := land[tree]; ok {
			land[tree]++
		} else {
			land[tree] = 1
		}
	}
	fmt.Printf("Result %v\n", land)
}
