# 📁 Duplicate File Finder
A command-line utility that walks a folder recursively, calculates SHA-256 hashes of all files, and groups files with identical hashes to find duplicates.

## 🚀 Features
* Walks nested directory trees recursively.
* Compares file contents using cryptographic hashing.
* Groups and outputs duplicate files together.

## 🛠️ Go Concepts Demonstrated
* **Directory Traversal**: Walking directory structures efficiently using `filepath.WalkDir`.
* **Cryptographic Hashing**: Computing hashes of file contents using `crypto/sha256` and `io.Copy`.
* **Data Grouping**: Mapping file hashes to path lists using `map[string][]string` for fast duplicates detection.
* **File System Inspection**: Validating directory existence and metadata using `os.Stat`.

## 📖 How to Run
Provide the directory path to check as an argument:
```bash
go run . /Users/khushii/my-target-folder
```
Example Output:
```bash
Scanning folder: /Users/khushii/my-target-folder

Duplicate Files:

Duplicate Group 1
--------------------
/Users/khushii/my-target-folder/image.png
/Users/khushii/my-target-folder/backup/image_copy.png
```
