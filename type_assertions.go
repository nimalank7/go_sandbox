package main

import "fmt"

func main() {
	var i interface{} = "hello"

	/*
		Type assertion which asserts that interface value i holds the concrete type string and assigns it to s.
	*/
	// Assigns the underlying value "hello" to s and asserts it is of type string
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)

	// hello, true
	fmt.Println(s, ok)

	f, ok := i.(float64)
	// 0, false
	fmt.Println(f, ok)

	// Panics
	f = i.(float64)
	fmt.Println(f)
}
