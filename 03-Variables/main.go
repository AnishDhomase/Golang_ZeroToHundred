package main

import "fmt"

const gst = 18
// rate := 10 //Error

func main() {
	var a int
	a = 1
	var b string = "Anish"
	var c bool = true
	var d float32 = 1.1

	var e = "ard" //Type Inference

	k := 10 //Short Declaration

	const l = "anish" //Constant
	const m string = "anish" //Constant
	// l = "anish" //Error
	const (
		port = 8080
		host = "localhost"
	) //Constant Block
	
	fmt.Println(a, b, c, d, e, k, l, m, port, host)
}