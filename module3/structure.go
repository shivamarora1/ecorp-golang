//Struct in Golang
package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	//Defining structure
	type person struct {
		Name   string
		Age    int
		Weight int
	}

	//Creating a variable of structure type
	var p1 person
	p1.Name = "Rob"
	p1.Age = 23
	p1.Weight = 67

	var arr []int
	arr = append(arr, 23)
	arr = append(arr, 34)
	arr = append(arr, 90)

	b, err := json.Marshal(arr)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(b))
	}

}
