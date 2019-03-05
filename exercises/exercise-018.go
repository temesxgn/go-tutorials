package main

import "fmt"

/**
Using a COMPOSITE LITERAL:
create a SLICE of TYPE int
assign 10 VALUES
Range over the slice and print the values out.
Using format printing
print out the TYPE of the slice

*/

func main() {
	x := []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512}

	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}
