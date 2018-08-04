package main

//imports are at the file level, not package level
//meaning you must import "fmt" in each file, even if it is in the same package
import "fmt"

//x not declared in a block, { }, so it is at package level scope
var x = 10

func main() {
	//you can access x here since it is package level
	fmt.Println(x)
	//y is declared inside a block, { }, so it is only accessible in func main
	var y = "ten"
	fmt.Println(y)
} //scope of y ends here

//func foo starts with a lower case letter, it is not visible outside the package
//a func call "Foo" would be visible outside package main
func foo() {
	//can also access x here
	fmt.Println(x)
	//this line causes compilation issues because it doesnt know y
	//fmt.Println(y)
}
