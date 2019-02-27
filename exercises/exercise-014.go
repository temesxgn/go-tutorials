package main

import "fmt"

/**
* Create a for loop using this syntax
* for condition { } & for { }
* Have it print out the years you have been alive.
 */
func main() {
	year := 1991
	for year <= 2019 {
		fmt.Println(year)
		year++
	}

	year = 1991

	for {
		if year > 2019 {
			break
		}

		fmt.Println(year)
		year++
	}
}
