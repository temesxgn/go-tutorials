package main

import "fmt"

/**
Using the code from the previous example,
delete a record from your map.
Now print the map out using the “range” loop

*/

func main() {
	m := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}

	delete(m, "no_dr")

	for k, v := range m {
		fmt.Println("Record ", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}

}
