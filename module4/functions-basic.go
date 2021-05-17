package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	a := 42
	b := 13
	fmt.Println(add(a, b))
}
