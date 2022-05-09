package main

import "fmt"

type Vehicle struct {
	Doors int
	Color string
}

type Truck struct {
	Vehicle   Vehicle
	FourWheel bool
}

type Sedan struct {
	Vehicle Vehicle
	Luxury  bool
}

func main() {
	t := Truck{
		Vehicle:   Vehicle{Doors: 4, Color: "Yellow"},
		FourWheel: true,
	}
	s := Sedan{
		Vehicle: Vehicle{Doors: 4, Color: "Blue"},
		Luxury:  true,
	}
	fmt.Printf("%+v\n", t)
	fmt.Printf("%+v\n", s)
	fmt.Println(t.Vehicle.Doors)
	fmt.Println(s.Vehicle.Doors)
}
