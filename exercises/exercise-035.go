package main

import "fmt"

/**
Build and use an anonymous func
*/

func main() {
	func() {
		fmt.Println("Printing from func!")
	}()
}
