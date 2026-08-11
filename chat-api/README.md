# 💬 Real-Time Chat API
A minimal WebSocket-based chat server demonstrating goroutines, channels, and concurrent connection management in Go.

## 🚀 Features
* **Live messaging**: Connect via WebSocket and broadcast messages to every other connected client, instantly.
* **JWT auth**: Connections are authenticated via a token passed as a query parameter (browsers can't set custom headers on WebSocket handshakes).
* **Structured messages**: Every broadcast message carries the sender's `user_id` alongside the content, not just raw text.
* **In-memory only**: No database — messages exist only while the server runs, keeping the project small and focused on concurrency.

## 🛠️ Go Concepts Demonstrated
* **Goroutines**: Each connected client runs two independent goroutines (`readPump` / `writePump`) — one per direction of traffic — started with the `go` keyword.
* **Channels**: All communication between goroutines happens through typed channels (`chan []byte`), never through directly shared variables, avoiding race conditions.
* **The Hub Pattern**: A single goroutine (`Hub.Run`) owns the set of connected clients and is the *only* code that ever reads or modifies that set — every other goroutine talks to it exclusively through `register`/`unregister`/`broadcast` channels.
* **`select` over multiple channels**: The Hub's event loop uses `select { case ... }` to react to whichever channel (register, unregister, broadcast) has data ready first.
* **Buffered vs. unbuffered channels**: Each client's outgoing `Send` channel is buffered (`make(chan []byte, 256)`) so a slow reader doesn't block the Hub's broadcast loop; the Hub's internal channels are unbuffered by contrast.
* **`defer` for guaranteed cleanup**: `readPump` uses `defer` to guarantee a client is unregistered and its connection closed, no matter how the read loop exits.
* **JSON encoding (`json.Marshal`)**: Wrapping raw incoming bytes into a structured `ChatMessage{ UserID, Content }` before broadcasting — the reverse direction of the `ShouldBindJSON` decoding used in earlier projects.

## 📖 Setup

1. Copy `.env.example` to `.env` and set a JWT secret.
2. Install dependencies and run:
```bash
   go mod tidy
   go run ./cmd
```

## 📖 API

### Get a test token
```bash
GET /test-token/:userId
```
Issues a JWT for the given user ID. Stands in for real signup/login, which this small project intentionally skips to stay focused on WebSockets.

### Connect to chat
Upgrades to a WebSocket connection. Any text message sent is broadcast as JSON to every connected client:
```json
{ "user_id": 1, "content": "hello everyone" }
```

## 🧪 Manual Testing
A static `chat-test-client.html` file is included — open it directly in a browser (no server needed for the page itself), paste a token from `/test-token/:userId`, and click Connect. Open it in a second tab with a different token to see messages broadcast live between clients.

## 📦 Project Structure