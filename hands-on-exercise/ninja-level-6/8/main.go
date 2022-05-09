package main

import "fmt"

func main() {
	a := trop()
	a()
}

func trop() func() {
	return func() {
		fmt.Println("This is a function")
	}
}
