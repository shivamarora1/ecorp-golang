package main

import (
	"fmt"
)

func Vaccinate(peoples ...string) {
	for _, p := range peoples {
		fmt.Printf("%s Vaccinated\n", p)
	}
}

func main() {
	fmt.Println("Hey")                     //Hey
	fmt.Println("Hey", " Buddy")           //Hey Buddy
	fmt.Println("Hey", " Buddy ", " Carl") //Hey Buddy Carl
	fmt.Println()                          //
}
