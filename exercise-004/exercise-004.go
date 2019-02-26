package main

import "fmt"

type mytype int

var myTypeVar mytype

func main() {
	fmt.Println(myTypeVar)
	fmt.Printf("%T\n", myTypeVar)
	myTypeVar = 42
	fmt.Println(myTypeVar)
}
