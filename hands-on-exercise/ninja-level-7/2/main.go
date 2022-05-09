package main

import (
	"fmt"
)

type Person struct {
	Name string
}

func main() {
	p1 := Person{
		Name: "James Bond",
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}

func changeMe(p *Person) {
	p.Name = "Miss Moneypenny"
	// (*p).Name = "Miss Moneyp"
}
