package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, v := range x {
		fmt.Println(v)
	}
	fmt.Printf("%T", x)
}
