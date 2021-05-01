//MAP is of reference type
package main

import (
	"fmt"
)

func main() {

	var m1 map[int]string

	m1 = make(map[int]string)
	m1[0] = "Tania"
	m1[1] = "Sugam"
	m1[2] = "Bod"
	m1[3] = "Chris"
	m1[4] = "Dany"

	m2 := m1
	m2[2] = "Bob"
	m2[5] = "Rock"

	fmt.Printf("Map is %v\n", m1)
	fmt.Printf("Slice is %v\n", m2)

}
