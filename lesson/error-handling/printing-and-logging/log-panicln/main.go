package main

import (
	"fmt"
	"log"
	"os"
)

// NOTE: Package log implements a simple logging package ... writes to standard error and prints the date and time of each logged message ...
// NOTE: log.Panicln() is equivalent to log.Println() followed by a call to panic().

func main() {
	defer foo()
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Panicln(err)
	}
}

func foo() {
	fmt.Println("When os.Exit() is called, deferred functions don't run")
}