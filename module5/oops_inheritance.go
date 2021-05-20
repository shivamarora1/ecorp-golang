package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

// define a method in Human
func (h *Human) Greet() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

type Student struct {
	Human
	school string
}

func (s Student) Study() {
	fmt.Println("I am studying now")
}

type Employee struct {
	Human
	company string
}

func (e Employee) Work() {
	fmt.Println("I am working now")
}

func (e Employee) Greet() {
	fmt.Println("I am greeting professionaly")
}

func main() {
	sam := Employee{Human{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}
	mark := Student{Human{"Mark", 25, "222-222-YYYY"}, "MIT"}

	sam.Greet()
	mark.Greet()
	mark.Study()
}
