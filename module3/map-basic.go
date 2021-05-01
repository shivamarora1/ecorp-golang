package main

import (
	"fmt"
)

func main() {
	//Internal representation
	//{"key1":"Val1","key2":"Val2"}
	//Syntax map[keyType]valueType

	//Declaring a map
	var m1 map[string]int
	//By default map is nil
	if m1 == nil {
		fmt.Println("map is NIL")
	}

	//Initialising a map
	m1 = make(map[string]int)
	if m1 == nil {
		fmt.Println("NIL map")
	} else {
		fmt.Println("map is not NIL")
	}

	//Inserting values into map
	m1["v1"] = 100
	m1["v2"] = 200
	m1["v3"] = 150
	fmt.Println(m1)

	//Checking if key exists inside a map or not
	keyN := "v2"
	val, ok := m1[keyN]
	if ok {
		fmt.Printf("%s key exists in map and has %d values\n", keyN, val)
	} else {
		fmt.Printf("%s key does not exists in map\n", keyN)
	}

	//Deleting a key from map
	delete(m1, "v1")
	fmt.Println(m1)
}
