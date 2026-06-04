package main

import (
	"fmt"
	"time"
)

func infinite(name string) {
	for i := 0; true; i++ {
		fmt.Println(i, name)
		time.Sleep(time.Second * 1)
	}
}

// Surprisingly, the expected output doesn't appear.
// Why? Because the program's main function exits before the goroutines finish execution, leading to an incomplete run.
func main() {
	go infinite("vinod")
	go infinite("gyan")
}
