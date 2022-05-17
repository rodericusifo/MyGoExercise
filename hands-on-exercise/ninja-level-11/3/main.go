package main

import (
	"log"
)

type customErr struct {
	message string
}

func (c *customErr) Error() string {
	return c.message
}

func NewCustomError(message string) *customErr {
	return &customErr{
		message: message,
	}
}

func main() {
	err := NewCustomError("custom-error: This is an Error")
	foo(err)
}

func foo(e error) {
	log.Println(e)
}
