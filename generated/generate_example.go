//go:generate go run generate_script.go

package main

import "fmt"

/*
Running go generate in this directory will call go run_generate_script.go. This generates a zz_generated.go file
*/
func main() {
	fmt.Println("This is the main program.")
}
