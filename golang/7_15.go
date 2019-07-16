// Ahad Bawany
// Daily code 7/15/2019:
// Given an order id, implement functions to get an item from the log or add an item to the log
// For simplicitiy's sake, I've chosen to do this program using the built in slices

package main

import (
	"fmt"

)

// For simplicity's sake, the order ID is assumed to be an integer. This can easily be modified
func record(sli []int, n int )([]int){
	sli = append(sli, n)
	return sli
}

// Returns the ith LAST id
func get_last(sli []int, i int)(int){
	return sli[len(sli) - i]
}


func main(){
	sli := []int{}
	sli = record(sli, 30202)
	sli = record(sli, 4020)
	sli = record(sli, 5)
	sli = record(sli, 811)
	fmt.Println(get_last(sli, 2))
}
