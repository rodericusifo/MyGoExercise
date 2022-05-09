package main

import "fmt"

func main() {
	favSport := "Hiking"
	switch favSport {
	case "Hiking":
		fmt.Println("I Really Like Hiking")
	case "Swimming":
		fmt.Println("I Really Like Swimming")
	default:
		fmt.Println("I don't have favorite Sport")
	}
}
