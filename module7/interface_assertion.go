package main

import (
	"fmt"
)

func main() {
	var x interface{} = "100"

	// var s string = x.(string)
	// fmt.Println(s)

	s, ok := x.(string)
	fmt.Println(s, ok)

	n, ok := x.(int)
	fmt.Println(n, ok)

	n = x.(int)
}
