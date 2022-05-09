package main

import "fmt"

func main() {
	m := map[string][]string{
		"bond_james":      {`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": {`James Bond`, `Literature`, `Computer Science`},
		"no_dr":           {`Being evil`, `Ice cream`, `Sunsets`},
	}

	m["ifo"] = []string{`Gaming`, `Coding`, `Reading`}

	for k, v := range m {
		fmt.Println(k, v)
		for i, v := range v {
			fmt.Printf("\t%v %v\n", i, v)
		}
	}
}
