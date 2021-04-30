//Slice is of reference type

package main

import (
	"fmt"
)

func main() {

	var arr1 [3]int
	arr1[0] = 10
	arr1[1] = 20
	arr1[2] = 30

	arr2 := arr1
	arr2[2] = 45

	fmt.Printf("Arr 1 is %v\n", arr1)
	fmt.Printf("Arr 2 is %v\n", arr2)

	var arr3 []int
	arr3 = append(arr3, 25)
	arr3 = append(arr3, 90)
	arr3 = append(arr3, 13)

	var arr4 []int
	arr4 = arr3
	arr4[2] = 93

	fmt.Printf("Arr 3 is %v\n", arr3)
	fmt.Printf("Arr 4 is %v\n", arr4)

}
