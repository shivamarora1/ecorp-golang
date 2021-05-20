package main

import (
	"fmt"
)

type Human interface {
	Speak()
	Walk()
	Eat()
}

type Male struct {
}

func (m Male) Aggresive() {
	fmt.Print("Male is Aggresive")
}
func (m Male) Speak() {
	fmt.Print("Male is Speaking")
}
func (m Male) Walk() {
	fmt.Print("Male is Walking")
}
func (m Male) Eat() {
	fmt.Print("Male is Eating")
}

type Female struct {
}

func (m Female) Politeness() {
	fmt.Print("Female is Polite")
}
func (m Female) Speak() {
	fmt.Print("Female is Speaking")
}
func (m Female) Walk() {
	fmt.Print("Female is Walking")
}
func (m Female) Eat() {
	fmt.Print("Female is Eating")
}

type Animal struct {
}

func (a Animal) Walk() {
	fmt.Printf("Animal is Walking")
}
func (a Animal) Eat() {
	fmt.Printf("Animal is Eating")
}

func EnrollInUniversity(h Human) {
	h.Walk()
	fmt.Print(" in the university\n")

}

func main() {
	sam := Female{}
	EnrollInUniversity(sam)
}
