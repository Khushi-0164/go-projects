package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("go run . server")
		fmt.Println("go run . client")
		return
	}

	switch os.Args[1] {
	case "server":
		StartServer()
	case "client":
		StartClient()
	default:
		fmt.Println("Invalid option")
		fmt.Println("Use 'server' or 'client'")
	}
}
