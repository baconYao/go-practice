// create a func with the identifier foo that 
// 	takes in a variadic parameter of type int
// 	pass in a value of type []int into your func (unfurl the []int)
// 	returns the sum of all values of type int passed in
// create a func with the identifier bar that 
// 	takes in a parameter of type []int
// 	returns the sum of all values of type int passed in

package main

import (
	"fmt"
)

func main() {
	a1 := []int{1,2,3,4,5,6,7,8}
	n1 := foo(a1...)
	fmt.Println(n1)

	a2 := []int{1,2,3,4,5,6,7,8}
	n2 := foo(a2...)
	fmt.Println(n2)
}


func foo(a ...int) int {
	s := 0
	for _, v := range a {
		s += v
	}
	return s
}

func bar(a []int) int {
	s := 0
	for _, v := range a {
		s += v
	}
	return s
}