package main

import "fmt"

// Criar um array com duas posições de inteiros e armazenar em uma variavel a soma total da lista.
// A variavel deve ser imprimida no console.

func main() {
	lista := []int{20, 2}

	somaTotal := 0

	for i := range lista {
		somaTotal += lista[i]
	}

	fmt.Println(lista)
	fmt.Printf("Soma: %v", somaTotal)
}
