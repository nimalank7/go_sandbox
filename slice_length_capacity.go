package main

import (
	"fmt"
	"strings"
)

// Length of a slice is the number of elements it contains
// Capacity of a slice is the no. of elements from the 1st element in the slice until the last element of the underlying array
func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length
	// Length is 0 and capacity is 6
	s = s[:0]
	printSlice(s)

	// Length is now extended to 4 and capacity 6
	// [2, 3, 5, 7]
	s = s[:4]
	printSlice(s)

	// Drop its first two values so this is s[2:4] with [5, 7]
	// Length is 2
	// Capacity is 4 as we count the elements from 5 (first element in the slice)
	s = s[2:]
	printSlice(s)

	// Nil slice which has len = capacity = 0 and no underlying array
	var b []int
	fmt.Println(b, len(b), cap(b))
	if b == nil {
		fmt.Println("nil!")
	}

	// make function allocates a zero-ed array and a slice that refers to that array
	// [0, 0, 0, 0, 0]
	a := make([]int, 5)
	fmt.Println("a")
	printSlice(a)

	ticTacToeBoard()
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func ticTacToeBoard() {
	/*
		Creates a tic-tac-toe board
		This is a slice of a slice and is a slice of string slices (e.g. []string)
	*/

	// Create a tic-tac-toe board
	// Slice of a slice
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
