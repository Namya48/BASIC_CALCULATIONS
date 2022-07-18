package main

import (
	"fmt"
)

func add_values(a int, b int) int {
	var c = a + b
	return c
}
func main() {
	var a, b, c int
	fmt.Println("Enter Your 1st Number :\n")
	fmt.Scan(&a)
	fmt.Println("Enter Your 2nd Number :\n")
	fmt.Scan(&b)
	c = add_values(a, b)
	fmt.Print("Sum : ", c)
}
