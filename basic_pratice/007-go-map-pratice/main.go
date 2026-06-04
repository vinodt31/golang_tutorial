package main

import "fmt"

type User struct {
	Name  string
	Email string
}

func main() {
	// crate a map
	studentMarks := make(map[string]int)

	// add key value pair in map
	studentMarks["vinod"] = 90
	studentMarks["gyan tiwari"] = 80
	studentMarks["shyam tiwari"] = 70

	// print map
	println(studentMarks)

	// access value from map using key
	println("vinod marks: ", studentMarks["vinod"])

	// delete key value pair from map
	delete(studentMarks, "shyam tiwari")

	// print map after delete
	println(studentMarks)

	// check if key exist in map
	if marks, ok := studentMarks["shyam tiwari"]; ok {
		println("shyam tiwari marks: ", marks)
	} else {
		println("shyam tiwari not found in map")
	}

	// iterate over map
	for name, marks := range studentMarks {
		println(name, "marks: ", marks)
	}

	// map with struct value
	users := map[int]User{
		1: {
			Name:  "Vinod",
			Email: "vinod@example.com",
		},
		2: {
			Name:  "Rahul",
			Email: "rahul@example.com",
		},
	}

	fmt.Println(users[1].Name)

	// map with key and value both are struct

	for key, value := range users {
		fmt.Println(key, "=", value)
	}
}
