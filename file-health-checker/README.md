# 📂 Concurrent File Health Checker
A highly efficient CLI tool that checks the existence, size, and last modified timestamps of multiple files in parallel using concurrency.

## 🚀 Features
* Processes multiple file paths concurrently.
* Displays file status (`FOUND` / `NOT FOUND`), size, and modification timestamp.

## 🛠️ Go Concepts Demonstrated
* **Goroutines**: Initializing lightweight async threads using the `go` keyword.
* **WaitGroups**: Using `sync.WaitGroup` to orchestrate and wait for multiple async routines.
* **Safe Closure Variable Capture**: Passing loop variables directly to the goroutine function scope (`go func(p string) { ... }(path)`) to avoid closures indexing incorrect memory addresses.
* **File Metadata**: Utilizing `os.Stat` to retrieve file properties.
* **Modular Codebase Layout**: Creating a separate `checker` package and importing it into the main executable module.

## 📖 How to Run
Provide one or more file paths as arguments:
```bash
go run main.go file1.txt file2.png file3.pdf
```
Example Output:
```bash
----------
File: file1.txt
Status: FOUND
Size: 1042 bytes
Last Modified: 2026-06-19 12:00:00 +0000 UTC
```