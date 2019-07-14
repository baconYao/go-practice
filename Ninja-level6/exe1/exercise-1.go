// create a func with the identifier foo that returns an int
// create a func with the identifier bar that returns an int and a string
// call both funcs
// print out their results

package main

import (
	"fmt"
)

func main() {
	f := foo()
	fmt.Println(f)
	b1, b2 := bar()
	fmt.Println(b1, b2)
}

func foo() int {
	return 1
}

func bar() (int, string) {
	return 2, "bar"
}