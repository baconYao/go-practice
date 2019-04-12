// Take the code from the previous exercise, then store the values of type person in a map with the key of last name. Access each value in the map. Print out the values, ranging over the slice.

package main

import (
	"fmt"
)

type person struct {
	first                   string
	last                    string
	favoriteIceCreamFlavors []string
}

func main() {
	me := person{
		first: "Bacon",
		last:  "Yao",
		favoriteIceCreamFlavors: []string{
			"Chocolate",
			"Milk",
			"Martini",
		},
	}

	she := person{
		first: "Judy",
		last:  "Ling",
		favoriteIceCreamFlavors: []string{
			"Vanilla",
			"Milk",
			"Strawberry",
		},
	}

	m := map[string]person{
		me.last:  me,
		she.last: she,
	}

	fmt.Println(m)

	for _, v := range m {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, val := range v.favoriteIceCreamFlavors {
			fmt.Println(i, val)
		}
		fmt.Println("-------")
	}

}
