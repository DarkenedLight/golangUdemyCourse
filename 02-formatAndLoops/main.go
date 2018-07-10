package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 200; i++ {
		//d: decimal
		//b: binary
		//x: hex
		//q: ascii
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	}
}
