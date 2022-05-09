package main

import "fmt"

func main() {
	x := "ifo"
	defer hello(x)
	fmt.Println("My Name is", x)
}

func hello(name string) {
	fmt.Println("Hello", name)
}
