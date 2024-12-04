package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X float64
	Y float64
}

// Method with a receiver
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Function which is a value receiver
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/*
Pointer receiver. If we didn't have a pointer we'd be operating on a copy of the object. If we were
merely retrieving the value then (v Vertex) would be fine
*/
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
	fmt.Println(Abs(v))
}
