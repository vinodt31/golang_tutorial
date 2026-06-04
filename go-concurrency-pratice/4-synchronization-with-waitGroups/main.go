package main

import (
	"fmt"
	"sync"
)

func count(name string) {
	for i := 0; i < 5; i++ {
		fmt.Println(i, name)
	}
}

func main() {

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		count("vinod")
		wg.Done()
	}()

	wg.Wait()
}
