package main

import (
	"fmt"
	"math"
)

// NOTE: a NON-POINTER RECEIVER
// 			 works with values that are POINTERS or NON-POINTERS.
// NOTE: a POINTER RECEIVER
// 			 only works with values that are POINTERS.

type Square struct {
	Length float64
}

func (s Square) Area() float64 {
	return math.Pow(s.Length, 2)
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * (math.Pow(c.Radius, 2))
}

type IShape interface {
	Area() float64
}

func Info(s IShape) {
	switch s := s.(type) {
	case *Circle:
		fmt.Println("The Area of Circle is", s.Area())
	case Square, *Square:
		fmt.Println("The Area of Square is", s.Area())
	}
}

func main() {
	s1 := Square{
		Length: 10,
	}
	Info(s1)
	s2 := &Square{
		Length: 20,
	}
	Info(s2)

	c1 := &Circle{
		Radius: 7,
	}
	Info(c1)
}
