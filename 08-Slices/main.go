package main

import "fmt"

func main() {
	// Slices are like dynamic arrays
	var nums []int
	fmt.Println(nums) // []
	fmt.Println(nums == nil) // true
	fmt.Println(len(nums)) // 0
}