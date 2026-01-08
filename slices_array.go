package main

import "fmt"

func main() {
	lista := make([]int, 1)
	lista = append(lista, 2)
	lista = append(lista, 3)
	lista = append(lista, 4)
	lista = append(lista, 5)
	lista = append(lista, 6)
	lista = append(lista, 7)
	lista = append(lista, 8)
	lista = append(lista, 9)
	lista = append(lista, 10)
	fmt.Printf("%v\n\n", lista[3])

	// for
	//
	// somaTotal := 0
	// for i := 0; i < len(lista); i++ {
	// 	somaTotal += lista[i]
	// }
	// fmt.Printf("Soma total: %d\n", somaTotal)

	// lista = append(lista, 8)

	// listaDeStrings := []string{"golango", "chi", "gin"}

	// listaDeStrings = append(listaDeStrings, "gorm", "echo", "beego")

	// fmt.Println("Lista pos1: ", lista[0])
	// fmt.Println("Lista pos1: ", lista[1])
	// fmt.Println("Lista pos1: ", lista[3])
	// fmt.Println("Tamanho da lista: ", len(lista))

	// fmt.Println("Lista de strings: ", listaDeStrings)
	// fmt.Println("Tamanho da lista de strings:", len(listaDeStrings))
	//

	var listaToda = []int{2, 4, 10, 25, 34, 34, 23, 23, 23, 3}
	ultimoElemento := listaToda[len(listaToda)-1]
	segundaLista := listaToda[:3]
	terceiraLista := listaToda[4:]
	// var listaMenorQueCinco = make([]int, 0)

	// fmt.Println("Lista zerada", listaMenorQueCinco)
	// for i := 0; i < len(listaToda); i++ {
	// 	if listaToda[i] < 5 {
	// 		listaMenorQueCinco = append(listaMenorQueCinco, listaToda[i])

	// 		fmt.Println("Posição[", i, "]", listaMenorQueCinco)
	// 	}
	// }

	// fmt.Println("Lista menor que cinco:", listaMenorQueCinco)

	fmt.Println("Segunda lista:", segundaLista)
	fmt.Println("Terceira lista:", terceiraLista)
	fmt.Println("Ultimo elemento:", ultimoElemento)

	var array = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array:", array)

	var slice = []int{6, 7, 8, 9, 10}
	fmt.Println("Slice:", slice)
}
