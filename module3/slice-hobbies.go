package main

import (
	"fmt"
)

func main() {

	var hobbies []string
	var sName string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&sName)

	for i := 1; true; i++ {
		fmt.Printf("Enter your %d hobby:", i)
		var hobby string
		fmt.Scanln(&hobby)
		if hobby == "none" {
			break
		} else {
			hobbies = append(hobbies, hobby)
		}
	}
	fmt.Printf("%s has %d hobbies \n %v\n", sName, len(hobbies), hobbies)

}
