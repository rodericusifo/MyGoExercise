package main

import (
	"fmt"
)

// SYNTAX: func (r receiver) identifier(parameter(s)) (return(s)) { code }

func main() {
	// Unfurling
	xi := []int{2, 3, 4, 5, 6, 7, 8, 9}
	x := sum("", xi...)
	fmt.Println("The total is", x)

	// Pass Nothing
	y := sum("")
	fmt.Println("The total is", y)

	// Pass String and Variadic Parameter
	z := sum("james", xi...)
	fmt.Println("The total is", z)
}

// NOTE: Variadic Parameter should be the last and there's only one
func sum(s string, x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position", i, "we are now adding", v, "to the total which is now", sum)
	}
	return sum
}
