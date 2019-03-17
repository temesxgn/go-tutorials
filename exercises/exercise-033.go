package main

import "fmt"

/**
Create a user defined struct with
	the identifier “person”
	the fields:
		first
		last
		age
attach a method to type person with
	the identifier “speak”
	the method should have the person say their name and age
create a value of type person
call the method from the value of type person

*/

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("Hello, my name is %s %s and I am %v years old\n", p.first, p.last, p.age)
}

func main() {
	p := person{
		first: "Thomas",
		last:  "Hagos",
		age:   27,
	}

	p.speak()

}
