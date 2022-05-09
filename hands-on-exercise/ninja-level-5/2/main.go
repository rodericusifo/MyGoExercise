package main

import "fmt"

type Person struct {
	FirstName               string
	LastName                string
	FavoriteIceCreamFlavour []string
}

func main() {
	ifo := Person{
		FirstName:               "Ifo",
		LastName:                "Krista",
		FavoriteIceCreamFlavour: []string{"Strawberry", "Vanila"},
	}
	krista := Person{
		FirstName:               "Krista",
		LastName:                "Ifo",
		FavoriteIceCreamFlavour: []string{"Chocolate", "Blueberry"},
	}

	m := map[string]Person{
		ifo.LastName:    ifo,
		krista.LastName: krista,
	}

	for _, mItem := range m {
		fmt.Println(mItem.FirstName)
		fmt.Println(mItem.LastName)

		for i, v := range mItem.FavoriteIceCreamFlavour {
			fmt.Printf("\t%v %v\n", i, v)
		}
	}
}
