package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 100; i++ {
		fmt.Printf("dec : %d \t bin : %b \t hex : %#x\n", i, i, i)
	}
}
