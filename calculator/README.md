# 🧮 Calculator CLI
A simple command-line calculator built with Go that performs basic mathematical operations.

## 🚀 Features
* **Addition**
* **Subtraction**
* **Multiplication**
* **Division** (with safe division-by-zero check)

## 🛠️ Go Concepts Demonstrated
* Command-line argument parsing using `os.Args`.
* String-to-integer conversion with `strconv.Atoi`.
* Switch-case conditional routing.
* Custom error propagation and multiple return values `(int, error)`.

## 📖 How to Run
Run the calculator by providing the operation and two integer values:
```bash
go run main.go add 10 5      # Output: 15
go run main.go sub 10 5      # Output: 5
go run main.go mul 10 5      # Output: 50
go run main.go divide 10 5   # Output: 2
```
