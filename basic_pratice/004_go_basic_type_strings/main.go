package main

import (
	"fmt"
	"strings"
)

func main(){
	fname := "vinod"
	lname := "tiwari"
	fullname := fname + " " + lname

	fmt.Println(strings.ToUpper(fullname))
}