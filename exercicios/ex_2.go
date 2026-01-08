package main

import "fmt"

/*
Dado um slice com os itens "2, 8, 3, 10, 5, 4, 7, 9, 1,4" que vao de 1...10,
efetuar a soma em duas variaveis, os primeiros numeros de 1...5 e a segunda de 6...10.
Imprimir os resultados
*/

func main() {
	slice := []int{2, 8, 3, 10, 5, 4, 7, 9, 1, 4}

	soma1, soma2 := 0, 0

	for _, value := range slice {
		if value <= 5 {
			soma1 += value
		} else {
			soma2 += value
		}
	}

	fmt.Printf("SOma lista 1: %v\n", soma1)
	fmt.Printf("SOma lista 2: %v", soma2)

}
