package main

import "fmt"

func main() {
	x := expo(count)
	fmt.Println(x)
}

func count(x ...int) int {
	result := 1
	for _, v := range x {
		result *= v
	}
	return result
}

func expo(f func(x ...int) int) int {
	s := []int{1, 2, 3, 4, 5}
	return f(s...)
}
