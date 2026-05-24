package main

import "fmt"

func main(){
	// for range over slice example
	product := []string{"apple", "banana", "orange"}

	for i, v := range product{
		fmt.Println("index:", i, "value:", v)
	}

	// map example
	nameAge := map[string]int{
		"vinod": 30,
		"john": 25,
		"alice": 28,
	}

	for name, age := range nameAge {
		fmt.Println("name:", name, "age:", age)
	}

	fmt.Println("print map : ", nameAge)

	// delete a key from map
	delete(nameAge, "john")
	fmt.Println("after delete map : ", nameAge)

	// value and ok pattern
	age1, ok1 := nameAge["vinod"]
	age2, ok2 := nameAge["test"]

	fmt.Println("vinod : " , nameAge["vinod"], "test : " , nameAge["test"])
	fmt.Println("age1:", age1, "ok1:", ok1)
	fmt.Println("age2:", age2, "ok2:", ok2)

}