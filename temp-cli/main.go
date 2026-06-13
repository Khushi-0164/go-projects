package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go [c|f] value")
		fmt.Println("Example: go run main.go c 100")
		return
	}

	mode := os.Args[1]

	value, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Println("Invalid number!")
		return
	}

	if mode == "c" {
		f := (value * 9 / 5) + 32
		fmt.Printf("%.2f°C = %.2f°F\n", value, f)

	} else if mode == "f" {
		c := (value - 32) * 5 / 9
		fmt.Printf("%.2f°F = %.2f°C\n", value, c)

	} else {
		fmt.Println("Invalid mode. Use 'c' or 'f'")
	}
}
