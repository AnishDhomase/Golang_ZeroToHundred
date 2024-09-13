package main

import (
	"fmt"
	// "math/rand"
	"time"
)

func processNum(ch chan int) {
	// num := <-ch
	// fmt.Println("Processing ... ",num)

	for num := range ch {
		fmt.Println("Processing ... ",num)
		time.Sleep(time.Second * 1)
	}
}

func sum(a int, b int, ch chan int) {
	ch <- (a + b)
}

func longProcess(ch chan bool) {
	defer func() { ch <- true }()
	fmt.Println("Started a long process")
	time.Sleep(time.Second * 2)
	fmt.Println("Long process finished") 
}

func sendEmail(emailCh <-chan string, done chan<- bool) {
	// emailCh is a receive-only channel, so we can only receive data from the channel
	// done is a send-only channel, so we can only send data to the channel 
	defer func(){ done <- true }()
	for email := range emailCh {
		fmt.Println("Email sent to ... ", email)
		time.Sleep(time.Second * 1)
	} //Blocking bcoz of range
}


func main() {
	// Channels
	// Channels are used to communicate between goroutines
	// Channels are used to send and receive data which is blocking
	// Channels are typed
	// Channels are thread-safe
	// Channels are used to synchronize goroutines
	// Channels are used


	// Create a channel
	// ch := make(chan string) //Create a channel of data type string
	// ch <- "Hello"           //Send a string value to the channel //channels are blocking
	// msg := <-ch
	// fmt.Println(msg) 		//Receive a string value from the channel
	// Deadlock error
	// What is a deadlock?
	// A deadlock is a situation where a goroutine is waiting for another goroutine to finish
	// But the other goroutine is also waiting for the first goroutine to finish
	// So, both goroutines are waiting for each other
	// And the program is stuck
	// The program is not making any progress

	
	// Sending data to a channel
	// numCh := make(chan int)
	// go processNum(numCh)
	// numCh <- 10
	// // time.Sleep(time.Second * 1)
	// for {
	// 	numCh <- rand.Intn(100)
	// }


	// Receiving data from a channel
	// go sum(10, 20, numCh)
	// num := <-numCh
	// fmt.Println("Received ... ",num)


	// WeightGroup using channels to wait for a goroutine to finish
	// var ch = make(chan bool)
	// go longProcess(ch)
	// <- ch //It blocks the main goroutine until a value is received from the channel



	// Buffered channels
	// Non-blocking channels
	// done := make(chan bool)
	// emailCh := make(chan string, 100) //Create a buffered channel of data type string with a capacity of 3
	// // After the channel is full, the send operation will block
	// go sendEmail(emailCh, done)
	// // Send data to a buffered channel
	// for i := 1; i <= 10; i++ {
	// 	emailCh <- fmt.Sprintf("email%d@gmail.com", i) //Non-blocking
	// }
	// fmt.Println("Done sending emails")
	// close(emailCh) //Close the channel to signal that no more data will be sent to the channel
	// <-done //Wait for the goroutine to finish
	// Receive data from a buffered channel
	// fmt.Println(<-emailCh)
	// fmt.Println(<-emailCh)



	// Multiple channels listening
	ch1 := make(chan int)
	ch2 := make(chan string)

	go func() {
		ch1 <- 10
	}()
	go func() {
		ch2 <- "anish"
	}()

	for i:=0; i<2; i++ {
		select {
		case ch1Val := <-ch1:
			fmt.Println("Received from ch1 ... ", ch1Val)
		case ch2Val := <-ch2:
			fmt.Println("Received from ch2 ... ", ch2Val)
		}
	}
}