package main

import (
	"fmt"
	"os"
)

func main() {
	content := `package main

import "fmt"

func GeneratedFunction() {
	fmt.Println("This is a generated function!")
}
`

	err := os.WriteFile("zz_generated.go", []byte(content), 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("generated.go has been created.")
}
