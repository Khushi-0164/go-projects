# 📓 Notes CLI (JSON Database)
A command-line note-taking utility that manages text-based notes inside a local JSON database and supports keywords search.

## 🚀 Features
* **Create/Read/Update/Delete** notes.
* **Search**: Perform full-text search across note titles and content bodies using keywords.
* **Persistence**: Auto-saves states to `notes.json`.

## 🛠️ Go Concepts Demonstrated
* JSON serialization (`json.MarshalIndent`) and parsing (`json.Unmarshal`).
* Reading and writing local files via `os` package utility methods.
* String pattern matching using the `strings.Contains` method.
* Multi-field struct tags and slice element mutation.

## 📖 How to Run
* **Add a Note:**
  ```bash
  go run main.go add "Go Tip" "Use defer for resource cleanup"
  ```
* **List Notes:**
  ```bash
  go run main.go list
  ```
* **Search Notes:**
  ```bash
  go run main.go search "defer"
  ```
* **Delete Note:**
  ```bash
  go run main.go delete 1
  ```
