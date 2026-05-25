package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hello World")
	})

	fmt.Println("Server is running on http://localhost:5000")
	
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}