package main

import "fmt"

/**
Create and use an anonymous struct.
*/

func main() {

	t := struct {
		firstName string
	}{
		firstName: "Temesxgn",
	}

	fmt.Println(t.firstName)
}
