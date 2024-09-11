package main

import "fmt"

func main() {
	// Only for construct for loop

	// While loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i ++
	}

	// Infinite loop
	// for {}

	// For loop
	for i := 1; i <= 5; i++ {
		if(i == 2){
			continue
		}
		if(i == 4){
			break
		}
		fmt.Println(i)
	}

	// range loop 
	for i := range 5 {
		fmt.Println(i)
	} //Prints 0 - 4

}