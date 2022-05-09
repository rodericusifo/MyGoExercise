package main

import (
	"fmt"
)

type Person struct {
	First string
	Last  string
}

type SecretAgent struct {
	Person Person
	Ltk    bool
}

// SYNTAX: func (r receiver) identifier(parameters) (return(s)) { code }

func (s SecretAgent) Speak() {
	fmt.Println("I am", s.Person.First, s.Person.Last)
}

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

	fmt.Printf("%+v\n", sa1)
	fmt.Printf("%+v\n", sa2)
	sa1.Speak()
	sa2.Speak()
}
