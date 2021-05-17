package main

import "fmt"

// simple function to add 1 to a
func add1(x int) int {
	x = x + 1
	return x
}

func main() {
	x := 3

	fmt.Printf("x = %d", x) // should print "x = 3"

	x1 := add1(x) // call add1(x)

	fmt.Println("x+1 = ", x1)
	fmt.Println("x = ", x)
}
