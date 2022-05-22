package main

import (
	"fmt"

	"github.com/MyGoExercise/hands-on-exercise/ninja-level-12/1/dog"
)

func main() {
	fmt.Println("5 Years Human =", dog.Years(5), "Years Dog")
	fmt.Println("10 Years Human =", dog.Years(10), "Years Dog")
	fmt.Println("15 Years Human =", dog.Years(15), "Years Dog")
}
