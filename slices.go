package main

import "fmt"

func main() {

	// Creates a slice literal and builds the underlying array at the same time
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	q = append(q, 4, 5)

	/*
		Range allows us to iterate over a slice
		i is the index and v is the value


		If we wanted to discard either the index or the value we can use _
		for i, _ := range q
	*/
	for i, v := range q {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// Array
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	fmt.Println(names)
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	// Modify the underlying array
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	// Slice expression which is equivalent to a[0:4]
	c := a[:]
	fmt.Println(c)

	// Creates a slice literal of structs
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}
