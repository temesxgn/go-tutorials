package main

import "fmt"

/**
create a func with the identifier foo that returns an int
create a func with the identifier bar that returns an int and a string
call both funcs
print out their results
*/

func main() {
	f := foo()
	b, a := bar()

	fmt.Println(f)
	fmt.Println(b, a)
}

func foo() int {
	return 1
}

func bar() (int, string) {
	return 2, "I am a string"
}
