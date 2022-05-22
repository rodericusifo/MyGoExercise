package main

import (
	"fmt"
	"log"
	"os"
)

// NOTE: Package log implements a simple logging package ... writes to standard error and prints the date and time of each logged message ...
// NOTE: When log.Println() called. the log.SetOutput() will write down thw output to the corespondent file

func main() {
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Open("no-file.txt")
	if err != nil {
		log.Println(err)
	}
	defer f2.Close()

	fmt.Println("check the log.txt file in the directory")
}
