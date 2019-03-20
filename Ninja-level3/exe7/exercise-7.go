// Building on the previous hands-on exercise, create a program that uses “else if” and “else”.

package main

import (
	"fmt"
)

func main() {
	x := "Bacon Yao"

	if x == "Bacon yao" {
		fmt.Println("It's not 'Bacon yao'")
	} else if x == "bacon Yao" {
		fmt.Println("It's not 'bacon Yao'")
	} else {
		fmt.Println(x)
	}
}
