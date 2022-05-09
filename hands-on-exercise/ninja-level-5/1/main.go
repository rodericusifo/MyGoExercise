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
	
	fmt.Println(ifo.FirstName)
	fmt.Println(ifo.LastName)

	for i, v := range ifo.FavoriteIceCreamFlavour {
		fmt.Printf("\t%v %v\n", i,v)
	}

	fmt.Println(krista.FirstName)
	fmt.Println(krista.LastName)

	for i, v := range krista.FavoriteIceCreamFlavour {
		fmt.Printf("\t%v %v\n", i,v)
	}
}
