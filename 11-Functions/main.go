package main

import "fmt"

// Different argument types
func add(a int, b int) int {
	return a + b
}

// Same argument type
func add1(a, b int) int {
	return a + b
}

// Multiple return values
func getHoliDays() (string, string) {
	return "Diwali", "Holi"
}

// In go functions are first class citizens
// They can be passed as arguments to other functions
// They can be returned from other functions
// They can be assigned to variables
func calc(a int, b int, op func (int, int) int) int {
	return op(a, b)
}

func greet(name string) func() {
	return func() {
		fmt.Println("Hello", name)
	}
}

func main() {
	sum := add(10, 20) // 30
	fmt.Println(sum)

	h1, h2 := getHoliDays() 
	fmt.Println(h1, h2) // Diwali Holi

	// Anonymous function
	mul := func(a, b int) int { 
		return a * b
	}
	r := calc(10, 20, mul)
	fmt.Println(r) // 200

	// Returning a function
	greet("Anish")() // Hello Anish
	greetAnish := greet("Anish")
	greetAnish() // Hello Anish
}