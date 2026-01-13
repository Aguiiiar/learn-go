package main

import "fmt"

func main() {
	sliceInt := reserveWithGeneric([]int{1, 2, 3, 4})
	sliceString := reserveWithGeneric([]string{"Max", "Maria", "Matheus"})

	fmt.Println(sliceInt)
	fmt.Println(sliceString)
}

func reverse(slice []int) []int {
	newInt := make([]int, len(slice))

	for i := range slice {
		newInt[len(slice)-1-i] = slice[i]
	}

	return newInt
}

// type contrainType interface{
// 	int | string
// }

func reserveWithGeneric[T any](slice []T) []T {
	newSlice := make([]T, len(slice))

	for i := range slice {
		newSlice[len(slice)-1-i] = slice[i]
	}

	return newSlice

}
