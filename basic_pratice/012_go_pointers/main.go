package main

import "fmt"


func main(){
	// example 1 how to declare a pointer and assign value to it
	var num int = 10
	var ptr *int = &num

	fmt.Println("value of num: ", num)
	fmt.Println("address of num: ", &num)
	fmt.Println("value of ptr: ", ptr)
	fmt.Println("value at the address stored in ptr: ", *ptr)


	// example 2 how to pass pointer to a function and update the value of variable using pointer
	empSalary := 5000
	fmt.Println("employee salary before update: ", empSalary)
	updateEmployeeSalary(&empSalary)

	
	fmt.Println("employee salary address:", &empSalary)
    fmt.Println("employee salary after update:", empSalary)
}

func updateEmployeeSalary(salary *int){
	*salary = 7000
}
