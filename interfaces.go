package main

import "fmt"

/*
Interfaces can be seen as holding a value and a concrete type
(value, type)
Value = value of a specific concrete type (presumably an implementation)
Type = concrete type
*/

type ExInt interface {
	M()
}

type ExStr struct {
	S string
}

// Method on ExStr implements the M() method from our ExInt interface
func (t ExStr) M() {
	fmt.Println(t.S)
}

// Empty interface which specifies 0 methods and so may hold values of any type
interface{}

/*
Not a nil interface as it has a concrete type
*/
type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	/*
	   Here the value is ExStr{"hello"} and the concrete type is ExStr
	*/
	var e ExInt = ExStr{"hello"}
	e.M()

	// Defining a variable of the interface
	var a ExInt

	// Calling M() calls the implementation ExStr's M method
	a = ExStr{"hello"}
	a.M()

    var i I

	/*
		Here t isn't initialized so the value is nil but the concrete type is T
	*/
	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()

	var b interface{}
	b = 42
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
