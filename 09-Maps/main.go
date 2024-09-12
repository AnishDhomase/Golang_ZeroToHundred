package main

import (
	"fmt"
	"maps"
)

func main() {
	// Basic
	// mp := make(map[string]int)
	// mp["key1"] = 10
	// mp["key2"] = 20
	// fmt.Println(mp) // map[key1:10 key2:20]
	// fmt.Println(mp["key1"]) // 10
	// fmt.Println(mp["key3"]) // 0 (default value) if key is not present
	// fmt.Println(len(mp)) // 2
	// delete(mp, "key2") // Deleting key2m
	// clear(mp) // Clearing the map


	// Initialize map with values during declaration
	// mp := map[string]int{"key1": 10, "key2": 20}
	// fmt.Println(mp) // map[key1:10 key2:20]


	// Check if key is present
	// mp := map[string]int{"key1": 10, "key2": 20}
	// val, ok := mp["key1"]
	// fmt.Println(val, ok) // 10 true


	// Equality of maps
	mp1 := map[string]int{"key1": 10, "key2": 20}
	mp2 := map[string]int{"key1": 10, "key2": 20}
	fmt.Println(maps.Equal(mp1, mp2)) // true


}