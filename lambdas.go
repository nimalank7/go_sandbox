package main

import (
	"fmt"
	"math"
)

// Pass in our lambda function
func compute(fn func(float64, float64) float64) float64 {
	// Call the lambda function
	return fn(3, 4)
}

/*
Function closure which works the same as JavaScript.
Here we are returning a function that takes an int and returns an int.
*/
func adder() func(int) int {
	sum := 0

	// Inner function which references the sum variable.
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {

	/*
		Define our lambda function which takes in 2 float64 and returns a float64
	*/
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	// Pass in our lambda function to compute
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	/*
		Here we return our closures into pos and neg.
		Next we run both of them for 0 - 9
	*/
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
