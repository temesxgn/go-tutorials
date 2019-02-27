package main

import "fmt"

/**
Write a program that prints a number in decimal, binary, and hex
*/

func main() {
	num := 255

	fmt.Printf("%d\t%b\t%#x\n", num, num, num)

}
