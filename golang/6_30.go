// Ahad Bawany
// Daily code 6/30:
// Given a list of numbers and a number key, print whether any two numbers from the list add up
// to the key
package main

import(
	"fmt"
)

func OneSum(ma map[int]int, key int) {
	// this just grabs each value using a for loop
	for _, value := range ma {
		// key - value is equal to the number we need to add to a value to get the key
		// each index is equal to its value, so if an index exists, that value exists
		// This is a one-liner. Only activates if ok is true!
		if _, ok := ma[key - value]; ok{
			// if that value exists, we know that two numbers in the map added together can equal the key
			fmt.Println("True")
			return
		}
	}
}

func main() {
	// Create the map we'll be using later, each index is equal to its value
	m := map[int]int{
		10: 10,
		15: 15,
		3: 3,
		7: 7,
	}
	OneSum(m, 17)
}
