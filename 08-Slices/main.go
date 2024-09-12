package main

// import (
// 	"fmt"
// 	"slices"
// )

func main() {
	// Slices are like dynamic arrays

	// Basic
	// var nums []int
	// fmt.Println(nums) // []
	// fmt.Println(nums == nil) // true
	// fmt.Println(len(nums)) // 0


	// make and append functions
	// var arr = make([]int, 5, 10) // make([]T, len, cap) 
	// arr = append(arr, 1, 2, 3, 4, 5) // append(slice, elements)
	// fmt.Println(arr) // [0 0 0 0 0]


	// New Syntax for slices
	// slc := []int{0, 0, 0}
	// for i:=1; i<=5; i++ {
	// 	slc = append(slc, i)
	// }
	// slc[0] = 11
	// fmt.Println(slc) // [11 0 0 1 2 3 4 5]


	// Copying slices
	// slc1 := []int{1, 2, 3}
	// slc2 := make([]int, len(slc1))
	// copy(slc2, slc1)
	// fmt.Println(slc2) // [1 2 3]


	// Subslices
	// slc := []int{1, 2, 3, 4, 5}
	// fmt.Println(slc[1:3]) // [2 3] // [start:end] end is exclusive
	// fmt.Println(slc[:3]) // [1 2 3]
	// fmt.Println(slc[1:]) // [2 3 4 5]
	// fmt.Println(slc[:]) // [1 2 3 4 5]


	// camparing slices
	// slc1 := []int{1, 2, 3}
	// slc2 := []int{1, 2, 3}
	// fmt.Println(slices.Equal(slc1, slc2)) // true


	// 2D slices
	// slc := [][]int{{1, 2, 3}, {4, 5, 6}}
	
}