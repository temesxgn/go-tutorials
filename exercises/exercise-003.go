package main

import "fmt"

/**
Using the code from 002
1.  At the package level scope, assign the following values to the three variables
    a) x = 42
    b) y = "James Bong"
    c) z = true

2.  In func main
    a) use fmt.Sprintf to print all of the VALUES to one single string.
    ASSIGN the returned value of TYPE string using the short declaration
    operator to a variable with the IDENTIFIER "s" and print "s"

*/

var x = 42
var y = "James Bond"
var z = true

func main() {

	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)

	// Compiler assigned values is called zero value

}
