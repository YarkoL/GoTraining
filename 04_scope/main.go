package main

import (
	"fmt"
	"github.com/YarkoL/GoTraining/04_scope/vis"
)

var x int = 42

func main() {
	fmt.Println(x)
	foo()
	bar()
	bar()

	inc := wrapper()
	fmt.Println(inc())
	fmt.Println(inc())
	//fmt.Println(wrapper()) //prints the address
	//fmt.Println(inc) //same as above
	baz()
}

func foo() {
	fmt.Println(x) //package scope
	vis.PrintVar() //from import
	fmt.Println(increment())
	fmt.Println(x)
}

func increment() int {
	fmt.Println("in increment func block")
	x++
	return x
}

func bar() {
	fmt.Println("call bar")
	y := 0
	increment := func() int {
		y++
		return y
	}
	fmt.Println(increment())
	fmt.Println(increment())
	//fmt.Println(increment)
}

func wrapper() func() int { //returns a func that returns an int
	fmt.Println("call wrapper")
	y := 0
	return func() int {
		y++
		return y
	}
}

func baz() {
	fmt.Println("call baz")
	inc := wrapper()
	fmt.Println(inc()) //new scope so wrapper is reset
	fmt.Println(inc())
}
