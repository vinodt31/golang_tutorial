package main

import (
	"fmt"
	"time"
)

func main() {
	// The program counts "dog" forever and never gets to "cat".
	go infiniteCount("dog")
	infiniteCount("cat")
}

func infiniteCount(thing string) {
	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Second * 1)
	}
}
