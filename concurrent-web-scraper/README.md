# 🕸️ Concurrent Web Scraper & Health Checker
A CLI health check tool that scans a list of URLs from a text file concurrently, checking their availability and writing a detailed JSON report.

## 🚀 Features
* Reads targets from a local file (`urls.txt`).
* Inspects websites concurrently using a pool of workers.
* Generates a formatted JSON report (`report.json`).

## 🛠️ Go Concepts Demonstrated
* **Worker Pool Pattern**: Distributing tasks via input (`jobs`) and output (`results`) channels to a configured set of background workers.
* **Synchronization**: Using `sync.WaitGroup` to coordinate worker lifecycles and safely wait for all goroutines to finish.
* **File I/O**: Reading configurations using `bufio.Scanner` and writing json files using `os.WriteFile`.
* **Latency Benchmarking**: Measuring elapsed time for HTTP requests using the `time` package.

## 📖 How to Run
Ensure `urls.txt` contains a list of URLs (one per line) and execute the program:
```bash
go run .
```
Example Output:
```bash
Worker processing: https://google.com
Worker processing: https://github.com
{URL:https://google.com StatusCode:200 ResponseTime:130.45ms Error:}
{URL:https://github.com StatusCode:200 ResponseTime:245.18ms Error:}
Total Time: 250.12ms
```
