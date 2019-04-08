// Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice:

// 	"James", "Bond", "Shaken, not stirred"
// 	"Miss", "Moneypenny", "Helloooooo, James."

// Range over the records, then range over the data in each record.

package main

import (
	"fmt"
)

func main() {
	data1 := []string{"James", "Bond", "Shaken, not stirred"}
	data2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	fmt.Println(data1)
	fmt.Println(data2)

	x := [][]string{data1, data2}

	for i, xs := range x {
		fmt.Println("Record: ", i)
		for j, val := range xs {
			fmt.Printf("\t Index position: %v \t Value: \t %v \n", j, val)
		}
	}
}
