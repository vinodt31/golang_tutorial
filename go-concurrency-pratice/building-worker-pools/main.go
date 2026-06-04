package main

func main() {

	c := make(chan string, 5)

	c <- "hello vinod"               // This will not block because the channel has buffer space
	c <- "welcome to Go concurrency" // This will also not block

	/*
		msg1 := <-c
		fmt.Println("received data : ", msg1)

		msg2 := <-c
		fmt.Println("received data : ", msg2)
	*/
}
