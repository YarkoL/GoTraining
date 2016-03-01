package main

import (
	"fmt"
)

var e int

func main() {
	a := 2
	b := "jee"
	c := 3.14
	d := true
	
	fmt.Printf("%v  (%T)\n", a, a)
	fmt.Printf("%v  (%T)\n", b, b)
	fmt.Printf("%v  (%T)\n", c, c)
	fmt.Printf("%v  (%T)\n", d, d)
	
	fmt.Println()

	e = 42
	var f string
	var g float64
	var h bool
	
	fmt.Printf("%v  (%T)\n", e, e)
	fmt.Printf("%v  (%T)\n", f, f)
	fmt.Printf("%v  (%T)\n", g, g)
	fmt.Printf("%v  (%T)\n", h, h)
}
