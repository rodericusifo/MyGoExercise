package main

import (
	"fmt"
	"log"
	"os"
)

// NOTE: Package log implements a simple logging package ... writes to standard error and prints the date and time of each logged message ...
// NOTE: ... the Fatal functions call os.Exit(1) after writing the log message ...
// NOTE: log.Fatalln() is equivalent to log.Println() followed by a call to os.Exit(1).

func main() {
	defer foo()
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Fatalln(err)
	}
}

func foo() {
	fmt.Println("When os.Exit() is called, deferred functions don't run")
}