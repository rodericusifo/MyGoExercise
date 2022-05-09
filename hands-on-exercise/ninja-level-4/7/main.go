package main

import "fmt"

func main() {
	a1 := []string{"James", "Bond", "Shaken, not stirred"}
	a2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	ma := [][]string{a1, a2}

	for i, a := range ma {
		fmt.Println(i, a)
		for i, v := range a {
			fmt.Printf("\t%v %v\n", i, v)
		}
	}
}