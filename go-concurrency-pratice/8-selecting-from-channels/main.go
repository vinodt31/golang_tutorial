package main

import "time"

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {

		for {
			time.Sleep(500 * time.Millisecond)
			ch1 <- "vinod"
		}
	}()

	go func() {
		for {
			time.Sleep(1 * time.Second)
			ch2 <- "gyan tiwari"
		}
	}()

	for {
		select {
		case msg1 := <-ch1:
			println("Received from ch1: ", msg1)
		case msg2 := <-ch2:
			println("Received from ch2: ", msg2)
		default:
			println("No messages received, doing other work...")
			time.Sleep(500 * time.Millisecond) // Simulate doing other work
		}
	}
}
