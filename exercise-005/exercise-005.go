package main

import (
	"fmt"
)

type mytype int

var myTypeVar mytype
var y int

func main() {
	fmt.Println(myTypeVar)
	fmt.Printf("%T\n", myTypeVar)
	myTypeVar = 42
	fmt.Println(myTypeVar)

	y = int(myTypeVar)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
