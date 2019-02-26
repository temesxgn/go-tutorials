package main

import "fmt"

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
