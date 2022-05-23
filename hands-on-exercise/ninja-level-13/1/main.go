package starting_code

import (
	"fmt"

	"github.com/MyGoExercise/hands-on-exercise/ninja-level-13/1/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
	fmt.Println(dog.YearsTwo(20))
}
