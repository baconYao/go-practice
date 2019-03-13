// Create a for loop using this syntax
// 	for condition { }
// Have it print out the years you have been alive.

package main

import (
	"fmt"
)

func main() {
	x := 1
	for x <= 25 {
		fmt.Println(x)
		x++
	}
}
