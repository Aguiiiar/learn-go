package main

import "fmt"

func main() {
	text := "palavra"
	tamanho := len(text)

	i := 0
	for i < tamanho {

		if text[i] == 'r' {
			break
		}
		fmt.Println(string(text[i]))

		i++
	}

}

type Pessoa struct{}

func NewPessoa() *Pessoa {
	return &Pessoa{}
}
