package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", redirectHandler)
	http.HandleFunc("/shortner", shortenURLHandler)

	fmt.Println("Server running on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
