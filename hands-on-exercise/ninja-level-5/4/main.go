package main

import "fmt"

func main() {
	x := struct {
		Day  int
		Note string
	}{
		Day:  1,
		Note: "Day One",
	}
	fmt.Printf("%+v", x)
}
