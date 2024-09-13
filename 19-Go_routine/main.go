package main

import (
	"sync"
	// "time"
)

// func task(id int) {
// 	println("Task", id, "done")
// }
func task(id int, w *sync.WaitGroup) {
	defer w.Done() //Execute this statement when the function exits //Decrement the counter
	println("Task", id, "done")
}

func main() {
	// Goroutines
	// Goroutines are lightweight threads
	// Goroutines can run concurrently


	// Sequential execution
	// for i := 1; i <= 10; i++ {
	// 	task(i)
	// }

	// Concurrent execution
	// for i := 1; i <= 10; i++ {
	// 	go task(i)
	// }
	// The main function is also a goroutine
	// The main function will exit as soon as the last statement is executed
	// So, we need to add a sleep statement to wait for the goroutines to finish
	// time.Sleep(time.Second * 1)



	// WaitGroup
	// wait until all goroutines are finished
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1) //Increment the counter
		go task(i, &wg)
	}
	wg.Wait() //Wait until the counter is 0

}