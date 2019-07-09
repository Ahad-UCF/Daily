// Ahad Bawany
// Daily code 7/3/2019:
// Given a list of numbers, return the smallest positive integer that isn't part of the list

package main

import(
	"fmt"
)

func smallestNotPresent(a []int)(int){
	// Set b to 1 since that's the smallest positive value possible
	b := 1
	for i := 0; i < len(a); i++{
		// If the value within a is equal to our current b, that can't be the smallest possible value not present
		if a[i] == b{
			b++
		}
	}
	return b
}

func main(){
	a := []int {3, 4, -1, 1}
	fmt.Println(smallestNotPresent(a))
}
