package main

import (
	"fmt"
	"net/http"
)

// this function is of type http.HandleFunc Object
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	//tells the http pkg to handle all requests to the root("/") with handler
	http.HandleFunc("/", handler)
	fmt.Printf("Server listening at 8080 ...")

	//this function will block until the program is terminated
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Errorf("error occured: %v", err)
	}
}
