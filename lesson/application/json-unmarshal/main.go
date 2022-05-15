package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Age   int    `json:"age"`
}

func main() {
	s := `[{"first":"James","last":"Bond","age":32},{"first":"Miss","last":"Moneypenny","age":27}]`
	fmt.Println(s)

	bs := []byte(s)
	people := []Person{}

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", people)
}
