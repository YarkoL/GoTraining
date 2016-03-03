package main

import "fmt"

const (
	a = iota
	b // 1
	n = 42
	c = iota * 10 // 3 * 10
	p = "I am a constant string"
	q //same as above
)

const (
	 _ = iota //no use for zero
	KB = 1 << (iota * 10) //bitwise shift to the left by 10
	MB = 1 << (iota * 10) //bitwise shift to the left by 20
	
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(n)
	fmt.Println(c)
	fmt.Println(p)
	fmt.Println(q)
	fmt.Printf("%b\t %d\n", KB, KB)
	fmt.Printf("%b\t %d\n", MB, MB)
}
