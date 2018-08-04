package main

import "fmt"

/*
same as exercise 01 but using a func expression
*/

func main() {
	aFunc := func(x int) (int, bool) {
		return x / 2, x%2 == 0
	}
	var h, e = aFunc(10)
	fmt.Println(h, e)
}
