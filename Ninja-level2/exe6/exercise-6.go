// Using iota, create 4 constants for the NEXT 4 years. Print the constant values.

package main

import (
	"fmt"
)

const (
	_     = iota
	a int = (iota + 2019)
	b int = (iota + 2019)
	c int = (iota + 2019)
	d int = (iota + 2019)
)

func main() {
	fmt.Println(a, b, c, d)
}
