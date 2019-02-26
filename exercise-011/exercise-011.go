package main

import "fmt"

// Using iota, create 4 constants for the last 4 years and print them
const (
	a = 2016 + iota
	b
	c
	d
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
