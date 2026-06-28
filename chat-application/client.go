package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func StartClient() {
	// Connect to server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to server")

	// Read username
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter your username: ")
	scanner.Scan()
	username := scanner.Text()

	// Send username to server
	_, err = conn.Write([]byte(username))
	if err != nil {
		fmt.Println("Error sending username:", err)
		return
	}

	// Start receiving messages
	go receiveMessages(conn)

	fmt.Println("You can start chatting!")
	fmt.Println("Type your message and press Enter.")

	// Read chat messages
	for scanner.Scan() {
		message := scanner.Text()

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
