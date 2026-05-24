package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {

	u := User{Name: "Vinod", Age: 20}
	fmt.Println(u.Intro())

}

// val receiver means this method receives a copy of the User
func (u User) Intro() string {

	return fmt.Sprintf("Hi, I am %s", u.Name)
}
