package main

import "fmt"

type Example interface {
	hello()
}

type ClassA struct {
}

func (ClassA) hello() {
	fmt.Println("hello")
}

type ClassB struct {
}

func (ClassB) hello() {
	fmt.Println("hello")
}

// Programming to an interface
// Any implementation of e will satisfy this
func example(e Example) {
	e.hello()
}

func main() {
	a := ClassA{}
	example(a)
}
