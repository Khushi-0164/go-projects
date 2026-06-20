# 🚀 Go Projects Portfolio

A collection of practical Go (Golang) projects focused on backend development, concurrency, APIs, and software engineering fundamentals.

---

## 📁 Project Directory

Here is a categorized map of the projects in this repository:

### 1. 🧮 CLI & Basic Utilities
*   **[Calculator CLI](./calculator)**: Basic math operations with user argument validation and division-by-zero checks.
*   **[Number Guessing Game](./numGame)**: Interactive guess validation against a randomly generated secret seed.
*   **[Temperature Converter CLI](./temp-cli)**: Decimal conversions between Celsius and Fahrenheit.
*   **[Age Calculator](./Age-calculator)**: Exact age breakdown in years, months, and days using the standard `time` package.

### 2. 💾 Persistent Data & File Systems
*   **[Todo List CLI](./todo-CLI)**: Task CRUD management utilizing local JSON storage.
*   **[Note Taking CLI](./notes-cli-json)**: Text note manager featuring key-word searches across note headings.
*   **[Expense Tracker CLI](./expense-tracker)**: Spend logging featuring total spend aggregations and formatted outputs.
*   **[File Analyzer](./file-analyzer)**: Calculates lines, words, and characters inside a target text file.
*   **[Log Analyzer CLI](./log-analyzer)**: Filters and tallies logs line-by-line based on `INFO`, `WARN`, and `ERROR` classifications.

### 3. 🧩 Architecture, Concurrency & Security
*   **[Student Management System](./student-management)**: Modularized structure containing storage layers and services, mutating slice elements in-place via pointer receivers.
*   **[Concurrent File Health Checker](./file-health-checker)**: Checks multiple local file states concurrently using Goroutines and WaitGroups.
*   **[Concurrent URL Health Checker](./url-health-checker)**: Pings multiple web addresses concurrently, measuring latencies and catching timeouts.
*   **[Secure Password Hasher CLI](./password-hasher)**: Secure keyboard input masking using `golang.org/x/term` and constant-time password verification via `bcrypt`.

### 4. 🔌 Web Routing & REST APIs
*   **[Todo REST API (Standard Library)](./todo-api)**: A REST API built purely with standard libraries (`net/http`) featuring manual route parsing.
*   **[Todo REST API (Gin Framework)](./todo-gin)**: Migrated REST endpoints utilizing Gin for route queries, JSON parameter binding, and response builders.

---

## 🛠️ Tech Stack & Key Concepts Mastered

*   **Language**: Go (Golang)
*   **Concurrency**: Goroutines, `sync.WaitGroup`, Async Closures, and Outbound HTTP Client Timeouts.
*   **Structures & Pointers**: Receiver Methods, Value vs Pointer Receivers, Slice Modifications via Direct Address Offsets.
*   **Web Frameworks**: Gin (`gin.Default`, `ShouldBindJSON`, `c.Param`).
*   **Security & Encryption**: Bcrypt hashing (`golang.org/x/crypto/bcrypt`), terminal input masking (`golang.org/x/term`), constant-time string comparisons.
*   **Persistence**: JSON Encoding/Decoding (`json.MarshalIndent` / `json.Unmarshal`), File reads/writes (`os.ReadFile` / `os.WriteFile`).

---

## 📖 How to Run the Projects

Each project subdirectory contains its own standalone **`README.md`** file detailing run commands and specific arguments. To test a project, navigate to its folder and execute it:

```bash
cd <project-folder-name>
go run main.go [arguments...]