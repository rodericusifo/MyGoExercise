package main

import (
	"fmt"
)

func main() {
	x := 42
	if x == 40 {
		fmt.Println("our value was 40")
	} else {
		fmt.Println("our value was not 40")
	}

	x = 434
	if x == 40 {
		fmt.Println("our value was 40")
	} else if x == 41 {
		fmt.Println("our value was 41")
	} else {
		fmt.Println("our value was not 40 or 41")
	}
}
