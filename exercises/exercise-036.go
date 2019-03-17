package main

import "fmt"

/**
Assign a func to a variable, then call that func
*/

func foo() {
	fmt.Println("Printing from foo")
}

func main() {
	f := foo

	f()
}
