package main

import (
	"fmt"
)

func main() {
	var data = []float64{1, 2, 4, 6, 8, 9}
	var avg = average(data...)
	fmt.Println("average of", data, "is", avg)
}

func average(nums ...float64) float64 {
	var avg = 0.0
	for _, v := range nums {
		avg += v
	}

	return avg / float64(len(nums))
}

/*
-using the "..." when declaring parameters and passing in arguments
makes the function a variadic function
-this means that the function can take 0 or more arguments for that parameter
*/
