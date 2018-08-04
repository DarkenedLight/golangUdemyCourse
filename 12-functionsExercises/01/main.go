package main

import "fmt"

/*
create a functiont hat accepts an int and has two returns, one is the arguement
divided by 2 and the other is a bool saying if the arguement is even or not
*/

func half(x int) (int, bool) {
	return x / 2, x%2 == 0
}

func main() {
	h, e := half(5)
	fmt.Println(h, e)
}
