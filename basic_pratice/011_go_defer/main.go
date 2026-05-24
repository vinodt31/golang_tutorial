package main

import "fmt"

func main(){
	fmt.Println("This will be printed first")
	defer fmt.Println("This will be printed last")
	fmt.Println("This will be printed second")
	

	// Multiple defer Statements
	// LIFO order (Last In First Out)
	defer fmt.Println("1")
    defer fmt.Println("2")
    defer fmt.Println("3")

	fmt.Println("Test return: ", test())
	
}

func test() int {
    defer fmt.Println("defer")

    return 10
}