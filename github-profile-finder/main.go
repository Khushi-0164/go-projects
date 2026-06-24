package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter GitHub Username:")
		username, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("please enter Usename")
			return
		}
		username = strings.TrimSpace(username)
		if username == "" {
			fmt.Println("Username cannot be empty")
			continue
		}
		fmt.Println("Searching for:", username)
		FetchProfile(username)
		FetchRepositories(username)

		fmt.Print("\nSearch another user? (y/n): ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(strings.ToLower(choice))
		if choice != "y" && choice != "yes" {
			fmt.Print("Goodbye!")
			break
		}
	}
}
