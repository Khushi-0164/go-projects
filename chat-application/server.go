package main

import (
	"fmt"
	"net"
	"strings"
	"sync"
)

var (
	clients = make(map[net.Conn]string)
	mutex   sync.Mutex
)

func StartServer() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server running on :8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)

	// Read username
	n, err := conn.Read(buffer)
	if err != nil {
		return
	}

	username := strings.TrimSpace(string(buffer[:n]))

	// Store client
	mutex.Lock()
	clients[conn] = username
	mutex.Unlock()

	fmt.Printf("%s joined the chat\n", username)
	broadcast(fmt.Sprintf("%s joined the chat\n", username), conn)

	// Read chat messages
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Printf("%s left the chat\n", username)

			removeClient(conn)

			broadcast(fmt.Sprintf("%s left the chat\n", username), conn)
			return
		}

		message := fmt.Sprintf(
			"%s: %s\n",
			username,
			strings.TrimSpace(string(buffer[:n])),
		)

		fmt.Print(message)
		broadcast(message, conn)
	}
}

func broadcast(message string, sender net.Conn) {
	mutex.Lock()
	defer mutex.Unlock()

	for conn := range clients {
		if conn == sender {
			continue
		}

		_, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error sending message:", err)
		}
	}
}

func removeClient(conn net.Conn) {
	mutex.Lock()
	defer mutex.Unlock()

	delete(clients, conn)
}
