// Ahad Bawany
// Daily code 12/22
// Verify Collatz

package main

import(
	"fmt"
)

func collatzCheck(n int){
	if n == 1 {
		fmt.Println("We reached 1")
		return
	}

	if n %2 == 0 {
		collatzCheck(n/2)
	} else {
		collatzCheck(3*n+1)
	}
}

func main(){
	collatzCheck(100)
}
