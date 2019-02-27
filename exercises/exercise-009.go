package main

import "fmt"

/**
Write a program that:

1. Assigns an int to a variable
2. Print variable in decimal, binary, hex
3. Shift the bits over 1 position to the left and assign that to a variable
4. Print that variable in decimal, binary, and hex
*/

func main() {
	a := 42
	fmt.Printf("%d\t%b\t%#x\n", a, a, a)

	b := a << 1
	fmt.Printf("%d\t%b\t%#x\n", b, b, b)
}
