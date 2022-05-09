package main

import "fmt"

func main() {
	x := Foo()
	fmt.Println(x)

	y, z := Bar()
	fmt.Println(y, z)
}

func Foo() int {
	return 12
}

func Bar() (int, string) {
	return 12, "March"
}
