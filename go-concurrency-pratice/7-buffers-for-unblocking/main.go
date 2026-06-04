package main

import "fmt"

func main() {

	c := make(chan string, 2) // Create a buffered channel with capacity of 2

	c <- "hello vinod"               // This will not block because the channel has buffer space
	c <- "welcome to Go concurrency" // This will also not block

	// Now the channel is full, the next send will block until there is space
	// c <- "this will block" // Uncommenting this line will cause a deadlock

	// To avoid blocking, we can read from the channel
	msg1 := <-c
	fmt.Println("received data : ", msg1)

	msg2 := <-c
	fmt.Println("received data : ", msg2)
}
