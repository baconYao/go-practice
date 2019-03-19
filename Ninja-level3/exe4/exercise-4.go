// Create a for loop using this syntax
// 	for { }
// Have it print out the years you have been alive.

package main

import (
	"fmt"
)

func main() {
	i := 1
	for {
		if i > 25 {
			break
		}
		fmt.Println(i)
		i++
	}
}
