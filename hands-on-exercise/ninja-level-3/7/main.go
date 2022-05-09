package main

import "fmt"

func main() {
	a := 2
	if a == 1 {
		fmt.Println("The Number is equal to 1")
	} else if a == 2 {
		fmt.Println("The Number is equal to 2")
	} else {
		fmt.Println("The Number is not equal to 1 or 2")
	}
}