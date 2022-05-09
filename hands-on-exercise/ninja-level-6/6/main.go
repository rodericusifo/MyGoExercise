package main

import "fmt"

func main() {
	func(a, b int) {
		fmt.Println(a + b)
	}(1, 2)

	func() {
		fmt.Println("Hello World")
	}()
}
