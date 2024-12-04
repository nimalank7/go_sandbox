package main

import "fmt"

/*
type String interface {
	String() string
}
Anything that implements the String() method implements this interface
*/

type Person struct {
	Name string
	Age  int
}

// Implements the String() method allowing fmt.Sprintf to print out the struct
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
