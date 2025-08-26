package main

import (
	"fmt"
	"net/http"
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
    http.ListenAndServe(":8080", nil)
}