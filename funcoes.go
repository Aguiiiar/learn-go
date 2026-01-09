package main

import "fmt"

func main() {
	ImprimirMensagem("Hello, world!")
	ImprimirMensagem("Hello, world!!")

	message := FuncComRetorno("tetete func")

	fmt.Println(message)

	nome, idade := NRetornoNaFuncao()

	fmt.Println(nome, idade)

	peso, altura := RetornoNomeadoFuncao()
	fmt.Println(peso, altura)
}

func ImprimirMensagem(message string) {
	fmt.Println(message)
}

func FuncComRetorno(message string) string {
	return message
}

func NRetornoNaFuncao() (string, int) {
	return "Max", 22
}

func RetornoNomeadoFuncao() (peso float32, altura float32) {
	return 94.5, 1.81
}
