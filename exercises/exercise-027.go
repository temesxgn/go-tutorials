package main

import "fmt"

/**
Take the code from the previous exercise,
then store the values of type person in a map with the key of last name.
Access each value in the map. Print out the values, ranging over the slice.
*/

type person struct {
	firstName string
	lastName  string
	flavors   []string
}

func main() {
	p1 := person{
		firstName: "Thomas",
		lastName:  "Hagos",
		flavors:   []string{"Chocolate", "Cookies & Cream", "Rocky Road"},
	}

	p2 := person{
		firstName: "Hana",
		lastName:  "Haile",
		flavors:   []string{"Vanilla", "Chocolate", "Cookies & Cream"},
	}

	personMap := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for _, v := range personMap {
		fmt.Println(v.firstName)
		fmt.Println(v.lastName)
		for i, val := range v.flavors {
			fmt.Println(i, val)
		}
		fmt.Println("-------")
	}
}
