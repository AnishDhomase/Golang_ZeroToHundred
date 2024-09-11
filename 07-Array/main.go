package main

import "fmt"

func main() {

	// string Array default values are empty string
	var Weekdays [7]string
	Weekdays[0] = "Monday"
	Weekdays[1] = "Tuesday"
	fmt.Println(Weekdays)  // [Monday]
	fmt.Println(len(Weekdays)) // Length of array is 7
	fmt.Println(Weekdays[0])   // Element at index 0 is Monday

	// int Array default values
	var vals [5]int
	vals[0] = 1
	fmt.Println(vals)  // [1 0 0 0 0] default value is 0

	// bool Array default values are false

	// Initialize array with values
	nums := [5]int{1, 2, 3, 4, 5}
	fmt.Println(nums) // [1 2 3 4 5]

	// 2D Array
	nums2D := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(nums2D) // [[1 2 3] [4 5 6]]

}