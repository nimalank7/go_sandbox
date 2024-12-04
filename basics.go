package main

import (
	"fmt"
	"math"
)

// Return value of (string, string)
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	const Truth = true
	// No initializer so type has to be declared
	var i int
	i = 4
	fmt.Println(i)

	// If type is omitted then the variable will take the type of the initializer
	var c = "no!"

	// Short variabl declaration
	k := 3

	a, b := swap("hello", "world")
	fmt.Println(a, b, c, k)

	fmt.Println("Go rules?", Truth)

	if v := math.Pow(3, 2); v < 10 {
		fmt.Println(v)
	} else {
		fmt.Printf("%g >= %g\n", v, 10)
	}
}
