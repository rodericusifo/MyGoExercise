package main

import "fmt"

type Person struct {
	First string
	Last  string
	Age   int
}

func (p Person) Speak() {
	fmt.Println("Hello My Name is", p.First, p.Last, "and I am", p.Age, "years old")
}

func main() {
	ifo := Person {
		First: "Ifo",
		Last: "Krista",
		Age: 23,
	}
	ifo.Speak()
}
