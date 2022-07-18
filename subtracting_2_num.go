package main

import (
	"fmt"
)

func subtract_values(a int, b int) int {
	var c = a - b
	return c
}
func main() {
	var a, b, c int
	fmt.Println("Enter Your First Number :\n")
	fmt.Scan(&a)
	fmt.Println("Enter Your 2nd Number :\n")
	fmt.Scan(&b)
	c = subtract_values(a, b)
	fmt.Printf("Difference : %d", c)
}
