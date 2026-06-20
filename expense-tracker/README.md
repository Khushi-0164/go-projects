# 💳 Expense Tracker CLI
A command-line expense management application to track spending with automated date formatting and total expense calculations.

## 🚀 Features
* **Add Expense**: Log cost details, description title, and date.
* **List Expenses**: Output logs formatted as rows.
* **Delete Expense**: Delete items by ID.
* **Summary**: Calculate total expenditures and get transaction volume stats.
* **Database**: Local storage using `expenses.json`.

## 🛠️ Go Concepts Demonstrated
* Date-time layout formatting using Go's unique formatting reference: `time.Now().Format("2006-01-02")`.
* Converting strings to double-precision floats using `strconv.ParseFloat`.
* Aggregating values using standard loops and slices.
* Output structural alignments using formatted string identifiers.

## 📖 How to Run
* **Add an Expense:**
  ```bash
  go run main.go add "Coffee" 4.50
  ```
* **List Expenses:**
  ```bash
  go run main.go list
  ```
* **Get Spend Summary:**
  ```bash
  go run main.go summary
  ```
* **Delete Expense:**
  ```bash
  go run main.go delete 1
  ```