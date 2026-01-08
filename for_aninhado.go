package main

import "fmt"

func main() {
	for numBase := 1; numBase <= 10; numBase++ {
		for multiplicao := 1; multiplicao <= 10; multiplicao++ {
			fmt.Printf("%d x %d = %d\n", numBase, multiplicao, numBase*multiplicao)
		}
	}
}
