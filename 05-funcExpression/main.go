package main

import "fmt"

//this is a func whose return type is a func that returns an int
func wrapper() func() int {
	var x = 0
	return func() int {
		x++
		return x
	}
}

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

	//foo2 is assinged the anonymous fun that is defined by the return
	//statemnt in func wrapper
	var foo2 = wrapper()
	fmt.Println(foo2())
	fmt.Println(foo2())
}

//anonymous funcs are funcs without a name, like in a fun expression
//func expressions are where you assign a fun to a variable
