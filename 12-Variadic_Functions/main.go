package main

import "fmt"

func sum(nums ...int) int {
	// nums is a slice
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
func print(args ...interface{}) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}

func main() {
	// Functions can take variable number of arguments
	nums := []int {1, 2, 3, 4, 5}
	result1 := sum(1, 2, 3, 4, 5)
	result2 := sum(nums...)
	fmt.Println(result1, result2) // 15 15

	print(1, true, "abc", 1.23) // 1 true abc 1.23


}