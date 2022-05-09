package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type User struct {
	First   string   `json:"first"`
	Last    string   `json:"last"`
	Age     int      `json:"age"`
	Sayings []string `json:"sayings"`
}

func main() {
	u1 := User{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := User{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := User{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []User{u1, u2, u3}
	fmt.Printf("%+v\n", users)

	encoder := json.NewEncoder(os.Stdout)
	err := encoder.Encode(users)
	if err != nil {
		fmt.Println( err)
	}

	fmt.Println();fmt.Println("---------------");fmt.Println()

	s := `[{"first":"James","last":"Bond","age":32,"sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"first":"Miss","last":"Moneypenny","age":27,"sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"first":"M","last":"Hmmmm","age":54,"sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println(s)

	users = []User{}

	decoder := json.NewDecoder(strings.NewReader(s))
	err = decoder.Decode(&users)
	if err != nil {
		fmt.Println( err)
	}
	fmt.Printf("%+v\n", users)
}
