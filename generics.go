package main

import (
	"fmt"
)

type Number interface {
	int64 | float64
}

func sumOf[T Number](input []T) T {
	var inc T

	for _, val := range input {
		inc += val
	}
	return inc
}

func main() {

	fmt.Println(sumOf([]int64{0, 1, 2, 234, 34, 666}))
	fmt.Println(sumOf([]float64{1.0, 2.3, 4.7, 3.4, 8.9}))

}
