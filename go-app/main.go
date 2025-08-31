package main

import (
	"fmt"
	"net/http"
	"os"
)

func home(w http.ResponseWriter, r *http.Request) {
	// Render the home html page from static folder
	http.ServeFile(w, r, "static/abcd.html")
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Go Test!")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/test", test)
	fmt.Println("Server running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
		os.Exit(1)
	}
}