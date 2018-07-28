package main

import "fmt"

func hello() {
	fmt.Print("hello ")
}

func world() {
	fmt.Println("world")
}

func main() {
	//world()
	defer world()
	hello()
}

/*
-the "defer" keyword makes that line of code execute right before the function
returns

-in the above example, it allows us to make the function "world()" to be called
when function main is about to return after hitting the closing bracket "}"

-this is usefull for working with files, as you can put "defer file.close()"
right after opening it so its together, but it wont get called until the
function that it is in returns

-organization and readability
*/
