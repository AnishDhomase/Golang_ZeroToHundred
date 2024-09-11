package main

import "fmt"

func main() {
	age := 18
	if age <= 10 {
		fmt.Println("Go and Study first!")
	} else if age < 18 {
		fmt.Println("You are not eligible to vote!")
	} else {
		fmt.Println("You can vote!")
	}

	// Logical Operators
	// &&, ||, !

	// Declare a variable inside if
	if age := 20; age > 50 {
		fmt.Println("Old", age)
	} else if age > 30 {
		fmt.Println("Adult", age)
	} else {
		fmt.Println("Young", age)
	}

	// Go does not support ternary operator
}