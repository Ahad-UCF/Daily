// Ahad Bawany
// Daily code 7/9/2019:
// Given a time n and a function f, do function f after time n
package main

import(
	"fmt"
	"time"
)

// Example takes an input of n milliseconds to sleep for beore calling the func
// I don't believe golang supports generic funcs as inputs, so this is used instead
func example(n int) func() int{
	time.Sleep(time.Duration(n) * time.Millisecond)
	f := func() int{
		return 3
	}

	return f
}

func main(){
	ex := example(10000)
	fmt.Println(ex())
}
