package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	x := Foo(s1...)
	fmt.Println("The Total of", s1, "is", x)

	s2 := []int{6, 7, 8, 9, 10}
	y := Bar(s2)
	fmt.Println("The Total of", s2, "is", y)
}

func Foo(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func Bar(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
