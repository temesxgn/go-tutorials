package main

import "fmt"

/**
* Create a program that uses a switch statement
* with the switch expression specified as a variable
* of TYPE string with the IDENTIFIER 'favSport'
 */
func main() {

	favSport := "Basketball"

	switch favSport {
	case "Basketball":
		fallthrough
	case "Soccer":
		fmt.Println("My favorite sport is Soccer")
	case "Football":
		fmt.Println("MY FAVORITE SPORT IS FOOTBALL")
	}
}
