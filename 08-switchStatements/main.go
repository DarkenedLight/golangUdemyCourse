package main

import "fmt"

func main() {
	var lookForMe = "bob"

	switch lookForMe {
	case "you":
		fmt.Println("found you ")
	case "me":
		fmt.Println("found me")
	case "her":
		fmt.Println("found her")
	case "him":
		fmt.Println("found him")
	default:
		fmt.Println("were you looking for", lookForMe+"?")
	}
}

/*
-no break statements needed as fallthrough does not happen by default
-fallthrough can be enforced by using the "fallthrough" keyword at the end of
case
-the switch expression can be ommitted and then the first case that results to
true will be executed. ex: case len(lookForMe) == 2:
*/
