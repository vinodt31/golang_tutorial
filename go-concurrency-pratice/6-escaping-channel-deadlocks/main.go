package main

import (
	"fmt"
	"time"
)

func countWithChannle(name string, c chan string) {
	for i := 0; i < 5; i++ {
		c <- name
		fmt.Println(time.Millisecond * 500)
	}

	close(c) // Close after all sends are done
}

func main() {

	c := make(chan string)

	c <- "hello vinod" // This will cause a deadlock because there is no goroutine receiving from
	msg := <-c
	fmt.Println("received data : ", msg)

	/*
		go
		fatal error: all goroutines are asleep - deadlock!

		goroutine 1 [chan send]:
	*/
}
