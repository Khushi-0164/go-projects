# 📄 File Analyzer
A CLI application that reads a target text file and generates descriptive statistics including character, word, and line counts.

## 🚀 Features
* Displays raw file text output.
* Counts total characters.
* Counts lines.
* Counts words (handling double spaces and layout breaks).

## 🛠️ Go Concepts Demonstrated
* Reading files in memory using `os.ReadFile`.
* Type conversion from binary byte slice to Go `string`.
* Text parsing using helper functions: `strings.Split`, `strings.TrimSpace`, and `strings.Fields`.
* Error handling for missing or unreadable files.

## 📖 How to Run
Provide the path of the text file you wish to analyze:
```bash
go run . sample.txt
```
Example Output:
```bash
The Number of characters are: 245
Lines: 12
Words: 48
```
