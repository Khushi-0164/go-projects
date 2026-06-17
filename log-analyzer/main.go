package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <logfile>")
		return
	}
	fileName := os.Args[1]

	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	text := string(data)

	error := 0
	info := 0
	warn := 0
	fmt.Println("===== Error Logs =====")
	lines := strings.Split(strings.TrimSpace(text), "\n")
	for _, line := range lines {
		words := strings.Fields(line)

		if len(words) == 0 {
			continue
		}

		logType := words[0]
		switch logType {
		case "INFO":
			info++
		case "ERROR":
			fmt.Println(line)
			error++
		case "WARN":
			warn++
		}
	}

	fmt.Println("===== Log Summary =====")
	fmt.Println("INFO:", info)
	fmt.Println("WARN:", warn)
	fmt.Println("ERROR:", error)

	total := error + warn + info
	fmt.Println("Total Logs", total)
}
