// https://docs.google.com/document/d/1ckYpi6hcRkaBUEk975f54oGsHYHu7GhzOk7-nOrkNxo/edit#heading=h.5j71o7kp1ek6

package main

import (
	"fmt"
)

type myType int

var x myType
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Printf("%v\n", x)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
