package main

import "fmt"

/**
 Create your own type “person”
which will have an underlying type of “struct”
so that it can store the following data:
first name
last name
favorite ice cream flavors
Create two VALUES of TYPE person.
Print out the values, ranging over the elements in the slice which stores the favorite flavors.
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

	fmt.Println(p1)
	for i, flavor := range p1.flavors {
		fmt.Println(i, flavor)
	}

	fmt.Println(p2)
	for i, flavor := range p2.flavors {
		fmt.Println(i, flavor)
	}
}
