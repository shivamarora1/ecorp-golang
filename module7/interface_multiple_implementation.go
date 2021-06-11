package main

import (
	"fmt"
	"sort"
)

type animal interface {
	breathe()
	walk()
}

type mammal interface {
	feed()
}

type lion struct {
	age int
}

func (l lion) breathe() {
	fmt.Println("Lion breathes")
}
func (l lion) walk() {
	fmt.Println("Lion walk")
}
func (l lion) feed() {
	fmt.Println("Lion feeds young")
}

type butterfly struct {
}

func (b butterfly) breathe() {
	fmt.Println("Butterfly breathing")
}

func (b butterfly) walk() {
	fmt.Println("Butterfly Walking")
}

func main() {
	var a animal
	l := butterfly{}
	a = l
	a.breathe()
	a.walk()

	var m mammal
	m = l
	m.feed()
}
