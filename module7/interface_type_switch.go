package main

import (
	"fmt"
)

type IHuman interface {
	Walking()
	Dancing()
}

type IAnimal interface {
	Speak()
	FindPredator()
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
	fmt.Println("Speaking...")
}
func (m Dog) FindPredator() {
	fmt.Println("FInding Predator...")
}

func Identify(m interface{}) {
	switch v := m.(type) {
	case nil:
		fmt.Println("v is nil")
	case IAnimal:
		fmt.Println("v is Animal", v)
	case IHuman:
		fmt.Println("v is Human", v)
	default:
		fmt.Println("type unknown")
	}
}
func main() {
	// m := 4
	// var a *int
	Identify(nil)
}
