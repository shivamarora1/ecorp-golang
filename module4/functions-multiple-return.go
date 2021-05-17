package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, _ := swap("hello", "world")
	fmt.Println(a)
}
