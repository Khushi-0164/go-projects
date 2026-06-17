package main

import (
	"fmt"
	"os"
	"strings"
	//"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <filename>")
		return
	}
	fileName := os.Args[1]

	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	text := string(data)
	fmt.Println(text)

	char := len(text)
	fmt.Println("The Number of characters are:", char)

	lines := strings.Split(strings.TrimSpace(text), "\n")
	fmt.Println("Lines:", len(lines))

	words := strings.Fields(text)
	fmt.Println("Words:", len(words))
}
