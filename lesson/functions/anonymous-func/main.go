package main

import (
	"fmt"
)

func main() {
	Foo()

	func() {
		fmt.Println("Anonymous func ran")
	}()

	func(x int) {
		fmt.Println("The meaning of life:", x)
	}(42)
}

func Foo() {
	fmt.Println("Foo ran")
}
