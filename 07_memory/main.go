package main

import "fmt"

const metersToYards float64 = 1.09361

func main() {
	var meters float64
	getMeters("Enter meters swam", &meters)
	yards := meters * metersToYards
	fmt.Println(meters, " meters is ", yards, " yards.")
}

func getMeters(prompt string, pointer *float64) {
	fmt.Println(prompt)
	fmt.Scan(pointer)
}
