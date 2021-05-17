package main

import (
	"fmt"
)

func Recovers() {
	a := recover()
	if a != nil {
		fmt.Println(a)
	}
}
func B() {
	defer Recovers()

	fmt.Println("Inside Func B")
	panic("Error Occured")
	fmt.Println("Recovered Func B")
}
func A() {
	fmt.Println("Inside Func A")
	B()
	fmt.Println("Recovered Func A")

}

func main() {
	fmt.Println("Inside Main Program")
	A()
	fmt.Println("End of Main Program")
}
