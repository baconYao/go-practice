// Using the code from the previous example, delete a record from your map. Now print the map out using the “range” loop

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

	// I want to delete 'moneypenny_miss'
	// Step1. try to check the key exists in map
	if v, ok := m["moneypenny_miss"]; ok {
		fmt.Println("Value: ", v)
		// delete the key 'moneypenny_miss' in the map 'm'
		delete(m, "moneypenny_miss")
		fmt.Println(m)
	}

	for k, v := range m {
		fmt.Println("Key: ", k)
		for i, v2 := range v {
			fmt.Printf("\t%d\t%v\n", i, v2)
		}
	}
}
