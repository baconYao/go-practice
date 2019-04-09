// Using the code from the previous example, add a record to your map. Now print the map out using the “range” loop

package main

import (
	"fmt"
)

func main() {
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	m["new_key"] = []string{`nonono, no~`, `Ok, Don't worry`, `what!`}

	for k, v := range m {
		fmt.Println("Key: ", k)
		for i, v2 := range v {
			fmt.Printf("\t%d\t%v\n", i, v2)
		}
	}
}
