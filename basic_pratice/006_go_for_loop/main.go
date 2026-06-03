package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {

		fmt.Println(i)
	}

	// we can use for as below also
	for i := range 5 {

		fmt.Println(i)
	}

	// case practice

	day := 4

	switch day {
	case 1:
		fmt.Println("Sunday")
	case 2:
		fmt.Println("Monday")
	case 3:
		fmt.Println("Tuesday")
	case 4:
		fmt.Println("Wednesday")
	case 5:
		fmt.Println("Thursday")
	case 6:
		fmt.Println("Friday")
	case 7:
		fmt.Println("Saturday")
	default:
		fmt.Println("Invalid day")
	}

}
