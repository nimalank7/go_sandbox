package main

import "fmt"

func main() {
	deferExample()
	deferStacked()

	v := Vertex{1, 2}
	p := &v

	// Instead of (*p).X we can do p.X
	p.X = 1e9
	q := &Vertex{1, 2} // has type *Vertex
	v3 := Vertex{}     // Implicit type X:0 and Y:0
	fmt.Println(q.X)
	fmt.Println(v3)
}

func deferExample() {
	defer fmt.Println("world")

	fmt.Println("hello")
	pointerExample()
}

/*
Defer uses a stack (LIFO) so i = 4 is the last defer but the first out.
defer fmt.Println(4)
defer fmt.Println(3)
defer fmt.Println(2)
defer fmt.Println(1)
defer fmt.Println(0)
*/

func deferStacked() {
	fmt.Println("counting")

	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

func pointerExample() {
	i, j := 42, 2701

	p := &i         // Set a pointer variable pointing to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

func arrayFunction() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// Slices
	var r []int = primes[1:4]
	fmt.Println(r)

	/*
		This builds an underlying array composed of {2, 3, 5...} then creates a slice literal.
		[]int{2, 3 ...} is a slice literal as there is no mention of the size of the array
	*/
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	// Slice default as this is the same as q[0:6]
	t := q[:]
	fmt.Println(t)

	// Slice literal of an array of structs
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
