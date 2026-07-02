# 🔗 URL Shortener
A lightweight HTTP web service that shortens long URLs into 6-character alphanumeric short codes and handles redirection of shortened URLs back to the original destination.

## 🚀 Features
* Generates random 6-character alphanumeric short codes.
* Handles HTTP redirection (`302 Found`) for short codes.
* Utilizes in-memory storage for mapping codes to original URLs.

## 🛠️ Go Concepts Demonstrated
* **HTTP Routing**: Setting up standard web server routing with `net/http` handlers.
* **JSON Marshalling**: Handling JSON requests and response payloads using the `encoding/json` package.
* **String Utilities & Randomness**: Generating unique short codes with `math/rand` and character slicing.
* **Data Structures**: Managing active URL mappings using Go's built-in `map[string]string`.

## 📖 How to Run
Start the web server:
```bash
go run .
```
Example Output:
```bash
Server running on :8080
```

In another terminal, shorten a URL:
```bash
curl -X POST http://localhost:8080/shortner -d '{"url": "https://google.com"}'
```
Example Output:
```bash
{"short_url":"http://localhost:8080/a1b2c3"}
```