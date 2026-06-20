# 🪵 Log Analyzer CLI
A tool that reads a text log file, parses entries line-by-line, prints active errors, and calculates logging statistics categorized by log level.

## 🚀 Features
* Scans log files line-by-line.
* Categorizes log lines by level tags: `INFO`, `WARN`, and `ERROR`.
* Prints all `ERROR` log lines instantly to console.
* Provides a summary counts report.

## 🛠️ Go Concepts Demonstrated
* File systems reading.
* String splitting (`strings.Split`) and blank spaces trimming.
* Token parsing using `strings.Fields`.
* Loop control syntax (`continue`) for empty lines.

## 📖 How to Run
Provide the path of the log file as an argument:
```bash
go run . app.log
```
Example Output:
```bash
===== Error Logs =====
ERROR: Failed to connect to database at 12:04:15
ERROR: Uncaught exception in main handler at 12:05:00
===== Log Summary =====
INFO: 42
WARN: 7
ERROR: 2
Total Logs: 51
```
