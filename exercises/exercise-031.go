package main

import "fmt"

/**
create a func with the identifier foo that
	takes in a variadic parameter of type int
	pass in a value of type []int into your func (unfurl the []int)
	returns the sum of all values of type int passed in
create a func with the identifier bar that
takes in a parameter of type []int
returns the sum of all values of type int passed in
*/

func foo(i ...int) int {
	sum := 0

	for _, v := range i {
		sum += v
	}

	return sum
}

func bar(i []int) int {
	sum := 0

	for _, v := range i {
		sum += v
	}

	return sum
}

func main() {
	lst := []int{2, 4, 6, 8}
	f := foo(lst...)
	b := bar(lst)

	fmt.Println(f)
	fmt.Println(b)
}
