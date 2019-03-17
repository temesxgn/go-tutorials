package main

import "fmt"

/**
Create a value and assign it to a variable.
Print the address of that value.
*/

func main() {
	v := 42

	fmt.Println(v)
	fmt.Println(&v)
}
