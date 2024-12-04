package main

import "fmt"

// List represents a single-linked list that holds values of any type
type List[T any] struct {
	next *List[T]
	val  T
}

// Index returns the index of x in s, or -1 if not found.
// Here [T comparable] is the type so this function works on all T's that satisfy the comparable constraint
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}

/*
Here we have 2 cases (i) T = int and (ii) T = string
*/

func main() {
	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))

	intList := List[int]{val: 4}
	fmt.Println(intList.val)
}
