// Use the “defer” keyword to show that a deferred func runs after the func containing it exits.

package main

import (
	"fmt"
)

func main() {
	defer bar()	
	defer foo()
	// defer bar()
	fmt.Println("Hello, go!")
	fmt.Println("Goodbye, go!")
}

func foo() {
	defer func() {
		fmt.Println("foo DEFER ran")
	}()
	fmt.Println("foo ran")
}

func bar() {
	fmt.Println("I'm bar")
}