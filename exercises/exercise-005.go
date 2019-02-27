package main

import (
	"fmt"
)

/**
Building on the code from the previous example
1.  At the package level scope, using the "var" keyword, create a VARIABLE with the IDENTIFIER "y".
    The variable should be of the UNDERLYING TYPE of your custom TYPE "x"
2.  In func main
    a) now use CONVERSION to convert the TYPE of the VLAUES stored in "x" to the UNDERLYING TYPE
    b) Then use the short declaration operator to ASSIGN that value to "y"
*/

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
