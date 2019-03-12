// Create TYPED and UNTYPED constants. Print the values of the constants

package main

import (
	"fmt"
)

func main() {
	x := 20
	fmt.Printf("%d\t%b\t%#x\n", x, x, x)

	y := x << 1
	fmt.Printf("%d\t%b\t%#x\n", y, y, y)
}
