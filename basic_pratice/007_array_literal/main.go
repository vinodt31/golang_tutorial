package main

import "fmt"

func main(){
	// fix size array, it will not grow
	var products [3] string
	products[0] = "Laptop"
	products[1] = "Mobile"
	products[2] = "Tv"
	//products[3] = "Fridge" // This will cause an error because the index is out of range
	//007_array_literal/main.go:11:11: invalid argument: index 3 out of bounds [0:3]
	
	fmt.Println(products)

	// array literal example
	marks := [3] int{80, 90, 65}
	marks[1] = 95
	fmt.Println(marks)
	fmt.Println(len(marks))


}