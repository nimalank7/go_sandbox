package main

import "fmt"

/*
Nil interface has no value nor concrete type.
Here no concrete type is used nor is i initialized.
*/
type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	// Panics
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
