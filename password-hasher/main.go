package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err

}
func verifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func readSecureInput(prompt string) (string, error) {
	fmt.Print(prompt)

	fd := int(syscall.Stdin)
	passwordBytes, err := term.ReadPassword(fd)
	if err != nil {
		return "", err
	}
	fmt.Println() // Print a newline because ReadPassword doesn't print the enter keystroke's newline
	return strings.TrimSpace(string(passwordBytes)), nil
}
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  go run main.go hash       - Interactive password hashing")
		fmt.Println("  go run main.go verify     - Interactive password verification")
		return
	}
	command := os.Args[1]
	switch command {
	case "hash":
		password, err := readSecureInput("Enter password to hash: ")
		if err != nil {
			fmt.Println("Error reading password:", err)
			return
		}

		if len(password) == 0 {
			fmt.Println("Password cannot be empty!")
			return
		}
		hash, err := hashPassword(password)
		if err != nil {
			fmt.Println("Error hashing password:", err)
			return
		}
		fmt.Println("\nGenerated Bcrypt Hash:")
		fmt.Println(hash)
	case "verify":
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter the hash to verify: ")
		hashInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading hash:", err)
			return
		}
		hash := strings.TrimSpace(hashInput)
		password, err := readSecureInput("Enter password to check: ")
		if err != nil {
			fmt.Println("Error reading password:", err)
			return
		}

		if verifyPassword(password, hash) {
			fmt.Println("\n Success: Password MATCHES the hash!")
		} else {
			fmt.Println("\n Error: Password DOES NOT match the hash.")
		}

	case "default":
		fmt.Printf("Unknown command '%s'. Available commands: hash, verify\n", command)
	}
}
