package main

import "fmt"

func main() {
	// Iterating over a slice
	// nums := []int{11, 22, 33, 44, 55}
	// for ind, num := range nums {
	// 	fmt.Println(ind, "->", num)
	// }

	// Iterating over a map
	// mp := map[string]int{"key1": 10, "key2": 20}
	// for key, val := range mp {
	// 	fmt.Println(key, "->", val)
	// }

	// Iterating over a string
	str := "abcABC012"
	for ind, charCode := range str { // char is unicode code point (rune)
		fmt.Println(ind, "->", string(charCode), "->", charCode)
	}

}