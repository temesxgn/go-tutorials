package main

import "fmt"

/**
A “callback” is when we pass a func into a func as an argument. For this exercise,
pass a func into a func as an argument
*/

func foo() int {
	return 5
}

func bar(f func() int) int {
	return f()
}

func main() {
	fmt.Println(bar(foo))
}
