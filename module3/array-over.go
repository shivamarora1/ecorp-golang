//Track score in each over
package main

import (
	"fmt"
)

func main() {

	var over [6]string

	//Initialsed scores
	var sb1, sb2, sb3, sb4, sb5, sb6 string

	fmt.Println("Enter score in 1st ball:")
	fmt.Scanf("%s", &sb1)

	fmt.Println("Enter score in 2nd ball:")
	fmt.Scanf("%s", &sb2)

	fmt.Println("Enter score in 3rd ball:")
	fmt.Scanf("%s", &sb3)

	fmt.Println("Enter score in 4th ball:")
	fmt.Scanf("%s", &sb4)

	fmt.Println("Enter score in 5th ball:")
	fmt.Scanf("%s", &sb5)

	fmt.Println("Enter score in 6th ball:")
	fmt.Scanf("%s", &sb6)

	//Now assigning scores to over array
	over[0] = sb1
	over[1] = sb2
	over[2] = sb3
	over[3] = sb4
	over[4] = sb5
	over[5] = sb6

	//At end print this array
	fmt.Printf("%v", over)
}
