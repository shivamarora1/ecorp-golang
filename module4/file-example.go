package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Opening a file ")
	var file, err = os.OpenFile("app.txt", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer file.Close()

	_, err = file.Write([]byte("Heelo Budddy"))
	if err != nil {
		fmt.Println("Error in writting ", err.Error())
	}
}
