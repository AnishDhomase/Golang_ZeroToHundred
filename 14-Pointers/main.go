package main

import "fmt"

// If you want to modify the original value, you need to pass the reference of the variable

func passByValueAndModify(a int) {
	a = 20
}
func passByReferenceAndModify(addr *int) {
	*addr = 20 // dereferencing
}
func slicePassByReferenceAndModify(addr *[]int) {
	*addr = append(*addr, 40) // dereferencing
}


func main() {
	a := 10
	b := 10
	arr := []int{10, 20, 30}
	passByValueAndModify(a)
	passByReferenceAndModify(&b) // passing the address of b
	slicePassByReferenceAndModify(&arr) // passing the address of arr
	println(a) // 10
	println(b) // 20
	fmt.Println(arr) // [10 20 30 40]

}