# 🌡️ Temperature Converter CLI
A command-line tool that converts temperatures between Celsius and Fahrenheit.

## 🚀 Features
* Convert Celsius to Fahrenheit (`c`).
* Convert Fahrenheit to Celsius (`f`).
* Supports decimal precision input.

## 🛠️ Go Concepts Demonstrated
* Float parsing using `strconv.ParseFloat`.
* Command-line argument validation and flags.
* Output formatting using `fmt.Printf` for double-decimal precision (`%.2f`).

## 📖 How to Run
Provide the target mode (`c` or `f`) followed by the temperature value:
* **Celsius to Fahrenheit:**
  ```bash
  go run main.go c 37.5
  # Output: 37.50°C = 99.50°F
  ```
* **Fahrenheit to Celsius:**
  ```bash
  go run main.go f 98.6
  # Output: 98.60°F = 37.00°C
  ```