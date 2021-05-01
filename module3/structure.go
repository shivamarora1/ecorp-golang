//Struct in Golang
package main

import (
	"fmt"
)

func main() {

	//Defining structure
	type person struct {
		name   string
		age    int
		weight int
	}

	//Creating a variable of structure type
	var p1 person
	p1.name = "Rob"
	p1.age = 23
	p1.weight = 67

	fmt.Printf("Name is %s\nAge is %d\nWeight is %d", p1.name, p1.age, p1.weight)

}
