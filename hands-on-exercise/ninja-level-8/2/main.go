package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First   string   `json:"first"`
	Last    string   `json:"last"`
	Age     int      `json:"age"`
	Sayings []string `json:"sayings"`
}

func main() {
	s := `[{"first":"James","last":"Bond","age":32,"sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"first":"Miss","last":"Moneypenny","age":27,"sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"first":"M","last":"Hmmmm","age":54,"sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println(s)

	bs := []byte(s)
	people := []Person{}

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", people)

	for i, person := range people {
		fmt.Println("Person #", i)
		fmt.Println("\t", person.First, person.Last, person.Age)
		for _, saying := range person.Sayings {
			fmt.Println("\t\t", saying)
		}
	}
}
