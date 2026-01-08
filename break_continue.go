package main

import "fmt"

func main() {
	texto := "palavra"
	for i := 0; i < len(texto); i++ {

		// if texto[i] == 'r' {
		// 	continue
		// }
		//
		if texto[i] == 'r' {
			fmt.Println("Encontrou 'r'")
			break
		}

		fmt.Println(string(texto[i]))
	}
}
