package main

import "fmt"

func main() {

	// shorthand - declare and assign using ':=', type auto set
	a := 10
	b := "golang"
	c := true
	// %T prints the type of the supplied variable
	fmt.Printf("type: %T \n", a)
	fmt.Printf("type: %T \n", b)
	fmt.Printf("type: %T \n\n", c)

	// var declaration - var *name* *type* = *value*
	// can also omit "= *value*" to set it to the types "zero value"
	// while you can declare the type while assigning, golang reccomends
	// omitting type because it will be inferred
	var d = 2
	var e = "hey"
	var f bool
	// %v prints the value, or zero value if none is assigned
	fmt.Printf("value: %v \n", d)
	fmt.Printf("value: %v \n", e)
	fmt.Printf("value: %v \n", f)
}
