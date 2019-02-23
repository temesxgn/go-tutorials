package main

import "fmt"

var packageLevelVariable = 4

func main() {
	name := "Temesxgn"
	name = "Temesxgn Gebrehiwet"

	fmt.Println("Hello " + name)
	foo()

	for i := 0; i < 4; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	fmt.Println("Package level variable ", packageLevelVariable)
}

func foo() {
	fmt.Println("I'm in foo")
}
