package main

import "fmt"

func main() {

	// vazio
	mapVazio := make(map[string]string)

	fmt.Println(mapVazio)

	user := map[string]string{
		"name": "John",
		"age":  "30",
		"city": "New York",
	}

	// uso do range para interar itens
	for key, value := range user {
		fmt.Printf("%s: %s\n", key, value)
	}

	// deletar item do map:
	m := make(map[string]int)

	m["sp"] = 10000000
	m["cg"] = 1000
	m["rj"] = 100000

	delete(m, "rj")

	fmt.Println(m)
}
