# 💬 Console TCP Chat Application
A console-based chat room application that runs either as a TCP server or a client, allowing multiple clients to connect and exchange messages in real-time.

## 🚀 Features
* Supports running as either server or client from the same binary.
* Broadcasts client messages to all other active clients.
* Notifies users of joins and leaves.

## 🛠️ Go Concepts Demonstrated
* **TCP Sockets**: Establishing connection listener and client connections using `net.Listen` and `net.Dial`.
* **Concurrency**: Processing multiple concurrent clients and handling simultaneous read/write operations via goroutines.
* **State Locking**: Utilizing `sync.Mutex` to secure shared server state (connected clients map) and prevent race conditions.
* **Terminal I/O**: Reading user keyboard inputs from the command line using `bufio.NewScanner`.

## 📖 How to Run
Start the server in one terminal:
```bash
go run . server
```
Example Output:
```bash
Server running on :8080
```

Start one or more client instances in other terminals:
```bash
go run . client
```
Example Output:
```bash
Connected to server
Enter your username: Alice
You can start chatting!
Type your message and press Enter.
```
