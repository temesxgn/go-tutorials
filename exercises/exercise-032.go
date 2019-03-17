package main

import "fmt"

/**
Use the “defer” keyword to show that a deferred func runs after the func containing it exits
*/

func foo() {
	fmt.Println("Printing from foo")
}

func bar() {
	fmt.Println("Printing from bar")
}

func main() {
	defer foo()
	bar()
}
