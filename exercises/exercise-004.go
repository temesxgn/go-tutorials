package main

import "fmt"

/**
1. Create your own type. Have the underlying type be an int.
2. Create a VARIABLE of your new TYPE with the IDENTIFIER "x" using the "VAR" keyboard.
3. In func main
    a) print out the value of x
    b) print out the type of x
    c) assign 42 to the VARIABLE x using the = operator
    d) print out the value of the variable x
*/

type mytype int

var myTypeVar mytype

func main() {
	fmt.Println(myTypeVar)
	fmt.Printf("%T\n", myTypeVar)
	myTypeVar = 42
	fmt.Println(myTypeVar)
}
