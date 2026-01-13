package main

import "os"

func main() {
	_, err := os.Open("teste")
	if err != nil {
		panic(err)
	}
}
