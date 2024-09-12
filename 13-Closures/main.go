package main

// Closures: Functions remember variables that are outside their body
func counter() (func () int, func () int){
	cnt := 0
	return func () int {
		cnt ++;
		return cnt;
	}, func () int {
		cnt --;
		return cnt;
	}
}

func main() {
	increment, decrement := counter()
	println(increment()) // 1
	println(increment()) // 2
	println(decrement()) // 1
	println(decrement()) // 0
	println(decrement()) // -1
	
}