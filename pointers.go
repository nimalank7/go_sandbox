package main

import "fmt"

func main() {
	i, j := 42, 2701

	// Point to i
	p := &i

	// Read i through pointer
	fmt.Println(*p)

	// Set i through the pointer
	fmt.Println(i)

	// Point to j
	p = &j

	// Divide j through the pointer
	*p = *p / 37

	// See the new value of j
	fmt.Println(j)
}
