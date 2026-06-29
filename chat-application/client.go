package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func StartClient() {
	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to server")

	scanner := bufio.NewScanner(os.Stdin)

	// Read username
	fmt.Print("Enter your username: ")
	scanner.Scan()
	username := strings.TrimSpace(scanner.Text())

	if username == "" {
		fmt.Println("Username cannot be empty.")
		return
	}

	// Send username to the server
	_, err = conn.Write([]byte(username))
	if err != nil {
		fmt.Println("Error sending username:", err)
		return
	}

	// Goroutine to receive messages
	go receiveMessages(conn)

	fmt.Println("You can start chatting!")
	fmt.Println("Type your message and press Enter.")

	// Read and send messages
	for scanner.Scan() {
		message := strings.TrimSpace(scanner.Text())

		if message == "" {
			continue
		}

		_, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Input error:", err)
	}
}

func receiveMessages(conn net.Conn) {
	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("\nDisconnected from server.")
			return
		}

		fmt.Print(string(buffer[:n]))
	}
}
