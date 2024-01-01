package main

// go mod init generics-example
// go get golang.org/x/exp/constraints

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func MapValues[T constraints.Ordered](values []T, mapFunc func(T) T) []T {
	var newValues []T
	for _, value := range values {
		newValue := mapFunc(value)
		newValues = append(newValues, newValue)
	}
	return newValues
}

func main() {
	ints := []int{1, 2, 4, 8}
	result := MapValues(ints, func(v int) int {
		return v * v
	})

	fmt.Println(result)

}
