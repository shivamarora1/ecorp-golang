package main

import (
	"fmt"
)

func main() {
	//Syntax []T (Without any lenght)
	var a []int
	a = append(a, 3)
	a = append(a, 5)
	a = append(a, 9)
	fmt.Printf("%v", a)

}
