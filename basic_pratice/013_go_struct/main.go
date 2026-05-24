package main

import (
    "fmt"
    "encoding/json"
)


func main(){

	// First way to create struct
	user := User {
		Id: 1,
		Name: "vinod kumar tiwari",
		Age: 20,
	}

	fmt.Println("User: ", user)

	// Second way to create struct

	emp := Employee{
		Name: "Vinod",
		Email: "a@test.com",
	}

	fmt.Println("Employee: ", emp)

	// Struct Tags Very important in APIs.
	jsonData, _ := json.Marshal(user)

	fmt.Println(string(jsonData))

	// Anonymous Struct
	user2 := struct {
		Name string
		Age int
	}{
		Name: "Vinod",
		Age: 30,
	}
	fmt.Println("Anonymous user2: ", user2)

	// Struct Comparison
	a := User{Id:1}
	b := User{Id:1}

	fmt.Println(a == b)

	// Struct Embedding (Like Inheritance)
	emp2 := Employee2{
		Person: Person{
			Name: "Vinod",
		},
		Salary: 50000,
	}

	fmt.Println("Employee2: ", emp2)


}

type User struct {
	Id int
	Name string
	Age int
}

// Struct Tags Very important in APIs.
type Employee struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

type Person struct {
    Name string
}

type Employee2 struct {
    Person
    Salary int
}

