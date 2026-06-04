package main

import (
	"fmt"
	"time"
)

func countWithChannel(name string, c chan string) {
	for i := 0; i <= 5; i++ {
		c <- name
		time.Sleep(500 * time.Millisecond)
	}

	close(c) // Close after all sends are done
}

// Leveraging channels for communication.
func main() {

	c := make(chan string)

	go countWithChannel("vinod", c)

	// for recive first element from channel
	data := <-c
	fmt.Println("received data : ", data)

	//for recive all element from channel
	for mesg := range c {
		fmt.Println(mesg)
	}

}
