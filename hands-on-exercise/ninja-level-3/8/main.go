package main

import "fmt"

func main() {
	switch {
	case true:
		fmt.Println("Should be printed")
	case false:
		fmt.Println("Should not be printed")
	}
}
