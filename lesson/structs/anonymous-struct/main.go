package main

import (
	"fmt"
)

func main() {
	p1 := struct {
		First string
		Last  string
		Age   int
	}{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}

	fmt.Println(p1)
}
