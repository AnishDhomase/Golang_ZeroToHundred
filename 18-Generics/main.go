package main

import "fmt"

func printIntSlice(arr []int) {
	for _, val := range arr {
		println(val)
	}
}

// Generic function
func printSlice[T any](arr []T) {
	for _, val := range arr {
		println(val)
	}
} //any == interface{}, comparable

// Specific Types Generic Function
func printIntOrStr[T int | string](arr []T) {
	for _, val := range arr {
		println(val)
	}
}
func printEntries[T any, U int | string](arr []T, user U) {
	for _, val := range arr {
		println(val, "entry by user-id", user)
	}
}

type StackOfInt struct {
	elements []int
}
type Stack[T any] struct {
	elements []T
}

func main() {

	// // Generic function
	// nums := []int{10, 20, 30}
	// words := []string{"one", "two", "three"}
	// bools := []bool{true, false}

	// printIntSlice(nums)
	// printSlice(bools)
	// printSlice(words)
	// printIntOrStr(nums)
	// printIntOrStr(bools) // Error: cannot use bool as type int or string in argument to printIntOrStr

	// struct with generic type
	st1 := StackOfInt {elements: []int{1, 2, 3}}
	st2 := Stack[string] {elements: []string{"one", "two", "three"}}
	fmt.Println(st1)
	fmt.Println(st2)
	entries := []string{"Pen", "Car", "Phone"}
	printEntries(entries, "anishdhomase")
}