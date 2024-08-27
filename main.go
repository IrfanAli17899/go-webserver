package main

import (
	"fmt"
	"net/http"
)

func hello_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Request received", *r)
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		message := r.FormValue("message")
		fmt.Fprintf(w, "Thank you, %s! Your message: %s", name, message)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/api/hello", hello_handler)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
