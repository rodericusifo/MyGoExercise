package main

import (
	"fmt"
)

type Person struct {
	First string
	Last  string
}

func (p Person) Speak() {
	fmt.Println("I am", p.First, p.Last, " - the Person Speak")
}

type SecretAgent struct {
	Person Person
	Ltk    bool
}

func (s SecretAgent) Speak() {
	fmt.Println("I am", s.Person.First, s.Person.Last, " - the secretAgent Speak")
}

type IHuman interface {
	Speak()
}

func Bar(h IHuman) {
	switch h := h.(type) {
	case Person:
		fmt.Println("I was passed into barrrrrr", h.First)
	case SecretAgent:
		fmt.Println("I was passed into barrrrrr", h.Person.First)
	}
	fmt.Println("I was passed into Bar", h)
}

type hotdog int

func main() {
	sa1 := SecretAgent{
		Person: Person{
			"James",
			"Bond",
		},
		Ltk: true,
	}

	sa2 := SecretAgent{
		Person: Person{
			"Miss",
			"Moneypenny",
		},
		Ltk: true,
	}

	p1 := Person{
		First: "Dr.",
		Last:  "Yes",
	}

	fmt.Println(sa1)
	sa1.Speak()
	sa2.Speak()

	fmt.Println(p1)

	Bar(sa1)
	Bar(sa2)
	Bar(p1)

	// conversion
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
