package main

import "fmt"

func main() {
	var x = 0

	//this is called a func expression
	//which is where you assign a function to a variable
	var foo = func() int {
		x++
		return x
	}

	//golang doesn't like when you declare funcs inside of funcs
	// func foo() int {
	// 	x++
	// 	return x
	// }

	fmt.Println(foo())
}
