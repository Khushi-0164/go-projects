# 🌐 Concurrent URL Health Checker
A CLI diagnostic tool that pings multiple website URLs concurrently to check their availability, records HTTP status codes, and measures response latency.

## 🚀 Features
* Checks multiple website URLs in parallel.
* Measures and displays query latency (response duration).
* Reports HTTP status codes.

## 🛠️ Go Concepts Demonstrated
* **Concurrency**: Utilizing goroutines and `sync.WaitGroup` to execute HTTP calls in parallel.
* **HTTP Client**: Constructing custom clients (`http.Client`) with configurable connection timeouts.
* **Resource Management**: Utilizing `defer res.Body.Close()` to prevent network connection and socket leaks.
* **Latency Benchmarking**: Using the `time` package (`time.Now()` and `time.Since()`) to track time durations.

## 📖 How to Run
Provide one or more website URLs as arguments:
```bash
go run main.go https://google.com https://github.com https://invalid-site.xyz
```
Example Output:
```bash
https://google.com -> UP ( 200 ) [ 120.45ms ]
https://github.com -> UP ( 200 ) [ 245.18ms ]
https://invalid-site.xyz -> DOWN
```
