package main

import (
	"fmt"
)

type Person struct {
	First string
	Last  string
	Age   int
}

type SecretAgent struct {
	Person Person
	Ltk    bool
}

func main() {
	sa1 := SecretAgent{
		Person: Person{
			First: "James",
			Last:  "Bond",
			Age:   32,
		},
		Ltk: true,
	}

	p2 := Person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
	}

	fmt.Printf("%+v\n", sa1)
	fmt.Printf("%+v\n", p2)

	fmt.Println(sa1.Person.First, sa1.Person.Last, sa1.Person.Age, sa1.Ltk)
	fmt.Println(p2.First, p2.Last, p2.Age)
}
