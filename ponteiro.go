package main

import "fmt"

func main() {
	ponteiroX := 5
	ponteiroY := &ponteiroX
	*ponteiroY = 10
	fmt.Println("== Main ==")
	ImprimirValores(&ponteiroX, ponteiroY)
	fmt.Println(ponteiroX, *ponteiroY)
	fmt.Println(&ponteiroX, ponteiroY)
}

func ImprimirValores(ponteiroX *int, ponteiroY *int) {
	// fmt.Println("== ImprimirValores ==")

	*ponteiroX = 20
	// fmt.Println(ponteiroX, ponteiroY)
	// fmt.Println(ponteiroX)
}
