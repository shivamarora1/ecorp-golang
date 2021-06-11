package main

import (
	"fmt"
)

func myFunc(a interface{}) {
	fmt.Println(a)
}

func main() {
	var my_age string
	my_age = "Hello Panda"

	myFunc(my_age)
}
