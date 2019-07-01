// Ahad Bawany
// Daily code 7/1/2019:
// Given a list of numbers, return a list of numbers where each number is the result of multiplying
// all other numbers in the list
package main

import (
	"fmt"
)

func main() {
	sli := []int{
		3,
		2,
		1,
	}
	fmt.Println(multiplArr(sli))
}

func multiplArr(sli []int)([]int){
	// Generate a slice with the length of the first slice. This will hold our answers
	sli2 := make([]int, len(sli))

	// Initialize all our values to 1 in the answers array
	for i := 0; i < len(sli); i++{
		sli2[i] = 1
	}
	// Cycle through the answers array
	for i := 0; i < len(sli); i++{
		// Cycle through each element in the first array
		for j := 0; j < len(sli); j++{
			if (i != j){
				// Multiply each element EXCEPT for the index that's the same in both arrays
				sli2[i] *= sli[j]
			}
		}
	}
	return sli2
}
