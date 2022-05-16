package main

import (
	"log"
	"os"
)

// NOTE: Package log implements a simple logging package ... writes to standard error and prints the date and time of each logged message ...
// NOTE: log.Println() calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Println().

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Println(err)
	}
}

