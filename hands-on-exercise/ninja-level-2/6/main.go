package main

import "fmt"

const (
	one = 2017 + iota
	two
	three
	four
)

func main() {
	fmt.Println(one, two, three, four)
}
