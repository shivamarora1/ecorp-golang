package main

import (
	"fmt"
)

type IHuman interface {
	Walking()
	Dancing()
}

type Men struct {
}

func (m Men) Walking() {
	fmt.Println("Walking...")
}
func (m Men) Dancing() {
	fmt.Println("Dancing...")
}

type Dog struct {
}

func (m Dog) Speak() {
	fmt.Println("Barking...")
}

func HumanDance(i interface{}) {
	//	variable_name.(typ_to_check)
	res, ok := i.(IHuman)
	if ok {
		res.Dancing()
	} else {
		fmt.Println("Not a instance of HUMAN")
	}
}

func main() {
	m := Dog{}

	HumanDance(m)
}
