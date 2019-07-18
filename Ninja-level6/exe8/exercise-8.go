// Create a func which returns a func
// assign the returned func to a variable
// call the returned func

package main

import (
	"fmt"
)

func main() {
	x := rf()
	fmt.Println(x())
}

func rf() func() int {
	return func() int {
		fmt.Println("I'm return func")
		return 99
	}
}