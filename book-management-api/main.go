package main

import (
	"fmt"
	"net/http"
)

func main() {
	connectDB()
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/books", booksHandler)
	http.HandleFunc("/books?id:", booksHandler)
	fmt.Println("Server is running on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
