package main

import (
	"fmt"
)

func main() {
	x := Bar()
	fmt.Printf("%T\n", x)
	fmt.Println(x())
}

func Bar() func() int {
	return func() int {
		return 451
	}
}
