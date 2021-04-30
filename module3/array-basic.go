//Array Basic
package main

import (
	"fmt"
)

func main() {
	// Syntax [n]T where n is size of array and
	// T is type of variables that array can store

	var a [3]int
	var b [5]string
	var c [10]float64

	fmt.Printf("a is %v\n", a)
	fmt.Printf("b is %v\n", b)
	fmt.Printf("c is %v\n", c)

	// Accessing value of arrays

	/* Note array index always start
	with 0 and ends with total length -1*/
	a[0] = 2
	a[1] = 3
	a[2] = 4
	fmt.Printf("%v", a)

	//What will happed if try to access the index out of bound
	//fmt.Printf("%v", a[4])
}
