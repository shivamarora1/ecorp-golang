package main

import (
	"fmt"
	"studentDeleteme"
)

func main() {
	st := studentDeleteMe.Student{}
	st.School = "KVS"

	if err := st.SetName("Harish"); err == nil {
		if err := st.SetEmail("harish@gamil.com"); err == nil {
			st.SetGrade(10)
			fmt.Printf("Name: %s \nEmail: %s \nGrade: %d ", st.GetName(), st.GetEmail(), st.GetGrade())
		} else {
			fmt.Print("Email is not valid")
		}
	} else {
		fmt.Print("Name is not valid")
	}
}
