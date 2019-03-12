// Create a variable of type string using a raw string literal. Print it.

package main

import (
	"fmt"
)

func main() {
	a := ` here is a string
	and a dog
	with a cat
	"here is another string"
	oh! A woman!
	`

	fmt.Println(a)
}
