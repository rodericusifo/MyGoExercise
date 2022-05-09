package main

import "fmt"

var a int
var b float64

func main() {
	x := 42
	y := 2.5678
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	a = 42
	b = 2.5678
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
}