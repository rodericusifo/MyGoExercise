package main

import (
	"fmt"
)

// NOTE: defer runs when the block function almost reaches the end point,
//			 exatly at the "return",
//			 But it will not run when the block function not reaches the end point.
func main() {
	defer Foo()
	Bar()
}

func Foo() {
	fmt.Println("Foo")
}

func Bar() {
	fmt.Println("Bar")
}
