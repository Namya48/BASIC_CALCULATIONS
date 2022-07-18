package main

import (
	"fmt"
)

func multiply_values(a int, b int) int {
	var c = a * b
	return c
}
func main() {
	var a, b, c int
	fmt.Print("Enter Your First Number :\n")
	fmt.Scan(&a)
	fmt.Print("Enter Your Next number :\n")
	fmt.Scan(&b)
	c = multiply_values(a, b)
	fmt.Print("Product : %d \n", c)
}
