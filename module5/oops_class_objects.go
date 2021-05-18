package main

import (
	"fmt"
)

// Declaring a struct
// You can also think it like class
type Book struct {

	// defining struct variables/data
	name   string
	author string
	pages  int
}

// Adding a function with Book struct
func (b Book) printDetails() {

	fmt.Printf("Book %s was written by %s.", book.name, book.author)
	fmt.Printf("\nIt contains %d pages.\n", book.pages)
}

func main() {

	// Creating object of Book struct
	book1 := Book{"Monster Blood", "R.L.Stine", 131}
	book1.printDetails()

	book2 := Book{"Alchemist", "Paulo Coelho", 211}
	book2.printDetails()

}
