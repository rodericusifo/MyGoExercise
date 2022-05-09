package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) Speak() {
	fmt.Println("Hello My Name is", p.Name)
}

type IHuman interface{
	Speak()
}

func SaySomething(h IHuman) {
	h.Speak()
}

func main() {
	p1 := &Person{
		Name: "Person 1",
	}
	SaySomething(p1)
	p1.Speak()

	// ERROR:cannot use p2 (variable of type Person) as IHuman value in argument to SaySomething: Person does not implement IHuman (method Speak has pointer receiver)
	// p2 := Person{
	// 	Name: "Person 1",
	// }
	// SaySomething(p2)
	// p2.Speak()
}