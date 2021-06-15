package main

import (
	"fmt"
	"net/http"
)

func main() {
	website := "https://www.google123.com"
	if _, err := http.Get(website); err != nil {
		fmt.Printf("Error is %s\n", err.Error())
	} else {
		fmt.Printf("Valid\n")
	}
}
