package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

/*
To have an 'error' state we need error values which are achieved by implementing the Error() string method of the Error interface
*/

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	// Here run an error so it calls the MyError.Error() method
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
