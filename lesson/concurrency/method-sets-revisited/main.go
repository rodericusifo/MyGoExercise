package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

type IShape interface {
	Area() float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func Info(s IShape) {
	fmt.Println("area", s.Area())
}

func main() {
	c := Circle{5}
	// Info(c)
	fmt.Println(c.Area())
}
