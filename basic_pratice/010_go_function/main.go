package main

import "fmt"

func main(){
	fmt.Println(add(5, 10))
	sub, mul :=subAndMultiply(15, 10);
	fmt.Println("sub : ", sub, "mul : ", mul)

	//IIFE (Immediately Invoked Function Expression) in Go

	res := func (name string) string {
		return name
	}("John Doe")

	fmt.Println("printName : ", res)
}

// simple function for adding two numbers and return it
func add(num1 int, num2 int) int {
 return num1 + num2
}

// a function which multiple return values
func subAndMultiply(num1 int, num2 int)(int, int){
	return num1 - num2, num1 * num2
}

