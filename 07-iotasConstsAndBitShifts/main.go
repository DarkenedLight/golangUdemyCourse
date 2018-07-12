package main

import "fmt"

const (
	_  = iota             // 0
	kB = 1 << (iota * 10) //1 << (1*10)
	mB = 1 << (iota * 10) //1 << (2*10)
	gB = 1 << (iota * 10) //1 << (3*10)
	tB = 1 << (iota * 10) //1 << (4*10)
)

func main() {
	fmt.Println("binary\t\tdecimal")
	fmt.Printf("%b\t", kB)
	fmt.Printf("%d\n", kB)
	fmt.Printf("%b\t", mB)
	fmt.Printf("%d\n", mB)
	fmt.Printf("%b\t", gB)
	fmt.Printf("%d\n", gB)
	fmt.Printf("%b\t", tB)
	fmt.Printf("%d\n", tB)
}

/*
-iota: defined as the ninth letter in the greek alphabet,
or as an extremely small number

-if defining multiple iota's int he same const() block, they start at 0 and
each time another is defined it increments. they reset to 0 when defined int a
new const() block

-consts do not have to fully capitalized. they also follow the export/visible to
other package rule (first letter capitalized = exported/visible to other pckgs)

-if exported, golang intellisense will throw a green squiggly saying you should
put a comment above its declaration saying why the const exists

-<< is left bitwise shifting, just like in c++
*/
