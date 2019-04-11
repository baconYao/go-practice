// Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
// 	first name
// 	last name
// 	favorite ice cream flavors
// Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors.

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

	fmt.Println(me.first)
	fmt.Println(me.last)
	for i, v := range me.favoriteIceCreamFlavors {
		fmt.Println(i, v)
	}

	fmt.Println(she.first)
	fmt.Println(she.last)
	for i, v := range she.favoriteIceCreamFlavors {
		fmt.Println(i, v)
	}

}
