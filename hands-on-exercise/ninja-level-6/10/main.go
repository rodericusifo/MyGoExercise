package main

import "fmt"

func main() {
	y := Foo()
	z := Foo()
	fmt.Println(y())
	fmt.Println(z())
}

func Foo() func() int {
	x := []int{1, 2, 3, 4, 5}

	return func() int {
		result := 1
		for _, v := range x {
			result *= v
		}
		return result
	}
}
