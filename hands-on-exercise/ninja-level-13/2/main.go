package main

import (
	"fmt"

	"github.com/MyGoExercise/hands-on-exercise/ninja-level-13/2/quote"
	"github.com/MyGoExercise/hands-on-exercise/ninja-level-13/2/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
