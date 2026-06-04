package main

import (
	"fmt"
)

func main() {

	//--------------- go make function for slice -----------------
	fmt.Println("-------------------- go make function for 1. slice---------------------------------")
	// make function is used to create slice with specific length and capacity
	// make([]T, len, cap)
	nameList := make([]string, 3)

	nameList = append(nameList, "vinod")
	nameList = append(nameList, "gyan tiwari")
	nameList = append(nameList, "shyam tiwari")
	nameList = append(nameList, "rajan tiwari")
	fmt.Println(nameList)

	fmt.Println("length of element:", len(nameList))
	fmt.Println("capacity of element:", cap(nameList))

	// in both case you will get different output because in first case we are creating slice with length 3 and capacity 3
	// but in second case we are creating slice with length 0 and capacity 3 so when we append element in first case
	// it will create new slice with double capacity and copy all element from old slice to new slice but in second case
	// it will not create new slice because we have already created slice with capacity 3 so it will append element in same slice
	// and increase length of slice by 1

	// if you want to create slice with specific length and capacity then you can use make function
	nameList2 := make([]string, 0, 3)
	nameList2 = append(nameList2, "vinod")
	nameList2 = append(nameList2, "gyan tiwari")
	nameList2 = append(nameList2, "shyam tiwari")
	nameList2 = append(nameList2, "rajan tiwari")
	fmt.Println(nameList2)

	fmt.Println("length of element:", len(nameList2))
	fmt.Println("capacity of element:", cap(nameList2))

	// -------------------- make example with map---------------------------------
	fmt.Println("-------------------- make example with 2. map---------------------------------")
	// Maps can be declared with make in two ways:
	//Without initial capacity:
	//m := make(map[string]int)

	//With initial capacity:
	//m := make(map[string]int, 100)

	// make(map[T]T, len)
	studentMarks := make(map[string]int, 5)
	studentMarks["vinod"] = 90
	studentMarks["gyan tiwari"] = 80
	studentMarks["shyam tiwari"] = 70
	fmt.Println(studentMarks)
	fmt.Println("map length: ", len(studentMarks))

	fmt.Println("-------------------- make example with 3. Channels---------------------------------")

	//Channels can be declared in two ways to define how they handle data flow:

	//Unbuffered Channel (1 argument):
	//ch := make(chan int)

	//Buffered Channel (2 arguments):
	//ch := make(chan int, 10)

	// 1 - declare unbuffered channel
	ch1 := make(chan string)

	// 2- start go routine for background work
	go func() {
		fmt.Println("Working doing some work in background...")
		//time.Sleep(200 * time.Second) // Simulate work with sleep
		ch1 <- "vinod"
		ch1 <- "gyan tiwari"
		ch1 <- "shyam tiwari"
		ch1 <- "rajan tiwari"
		close(ch1) // Close the channel after sending all data

		fmt.Println("working sent data to channel")
	}()

	// 3. main goroutine will wait for data from channel
	data := <-ch1
	fmt.Println("main goroutine received data from channel: ", data)

	// Buffered Channel (With Capacity)

	ch := make(chan int, 4)

	// We can send 3 items without blocking, even without a receiver ready
	ch <- 10
	ch <- 20
	ch <- 30
	ch <- 50

	// If we tried to do `ch <- 40` here, the program would deadlock
	// because the buffer is full (capacity is 3)

	// Read the items out of the buffer
	fmt.Println(<-ch) // Prints: 10
	fmt.Println(<-ch) // Prints: 20
	fmt.Println(<-ch) // Prints: 30
	//fmt.Println(<-ch)

}
