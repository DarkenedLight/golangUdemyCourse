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
	//only way to have a func inside a func is with a func expression
	// func foo() int {
	// 	x++
	// 	return x
	// }

	fmt.Println(foo())

	fmt.Println("----------")

	//foo2 is assinged the anonymous fun that is defined by the return
	//statemnt in func wrapper
	var foo2 = wrapper()
	fmt.Println(foo2())
	fmt.Println(foo2())

	fmt.Println("----------")

	//the variable x incremented here is in a different scope that foo2 and is
	//reset to 0 when wrapper is called
	var foo3 = wrapper()
	fmt.Println(foo3())
	fmt.Println(foo3())

	fmt.Println("----------")

	//these both print 3 proving that they are incrementing independently
	fmt.Println(foo2())
	fmt.Println(foo3())
}

//anonymous funcs are funcs without a name, like in a fun expression
//func expressions are where you assign a fun to a variable
