package main

import (
	"fmt"
	"os"
)

func ShowText() {
	fmt.Println("finalizando manipulação do arquivo...")
}

func main() {
	file, err := os.Create("./settings.txt")
	file.Close()
	defer ShowText()

	if err != nil {
		panic(err)
	}

	_, err = file.Write([]byte("teste"))
	if err != nil {
		panic(err)
	}
}
