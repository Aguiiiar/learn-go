package main

import (
	"fmt"
	"os"
)

func ReadFile() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error:", r)
		}
	}()

	_, err := os.Open("./settings.text")

	if err != nil {
		panic("FileNotExist")
	}
}

func main() {
	ReadFile()
	fmt.Println("End")
}
