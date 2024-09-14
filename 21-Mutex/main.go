package main

import (
	"fmt"
	"sync"
)

type Post struct {
	views int
	mu sync.Mutex
}

func (P *Post) incViews(w *sync.WaitGroup) {
	defer w.Done()
	P.mu.Lock() //Lock the variable from other goroutines
	P.views ++
	P.mu.Unlock() //Unlock the variable for other goroutines
	// Use lock and unlock where variable is accessed by multiple goroutines 
}

func main() {
	// Mutex (Mutual Exclusion)
	// Mutex is used to handle the race condition
	// Race condition is a situation where two or more goroutines access the same variable

	p1 := Post{views: 0}
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go p1.incViews(&wg)
	}
	wg.Wait()
	fmt.Println(p1.views) 
	// when not go, 1000
	// when go and not mutex, varries Ex. 993, 998, 999, 1000 //due to race condition
	// when go and mutex, 1000

}