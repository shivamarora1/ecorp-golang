// Go program to illustrate the
// concept of the defer statement
package main

import "fmt"

// Functions
func mul(a1, a2 int) int {

	res := a1 * a2
	fmt.Println("Result: ", res)
	return 0
}

func show() {
	fmt.Println("Hello!, WORLD")
}

// Main function
func main() {

	mul(23, 45)
	defer mul(23, 56)
	show()
}
