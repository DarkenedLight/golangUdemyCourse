package main

import "fmt"

/*
write a function with one variadic parameter that finds the greatest number in
a list of numbers
*/

func greatest(list ...int) int {
	var great int
	for _, v := range list {
		if v > great {
			great = v
		}
	}
	return great
}

func main() {
	data := []int{72, 93, 129, 9, 238, 38, 29}
	fmt.Println(greatest(data...))
}
