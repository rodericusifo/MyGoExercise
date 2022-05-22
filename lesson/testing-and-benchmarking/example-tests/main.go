package main

import (
	"fmt"

	"github.com/MyGoExercise/lesson/testing-and-benchmarking/example-tests/acdc"
)

func main() {
	fmt.Println(acdc.Sum(2, 3))
	fmt.Println(acdc.Sum(2, 3, 4, 5, 6, 7, 8, 9))
}
