package main

import (
	"fmt"
)

func main() {
	x := 1

	for x < 10 {
		fmt.Println(x)
		x++
	}
	fmt.Println("done.")

	for {
		if x > 9 {
			break
		}
		fmt.Println(x)
		x++
	}
	fmt.Println("done.")
}
