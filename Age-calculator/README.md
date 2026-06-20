# 📅 Age Calculator
A CLI application that calculates a person's exact age in years, months, and days based on their birth date.

## 🚀 Features
* Dynamic date calculations using the standard library.
* Outputs age breakdowns in three units: years, months, and total days lived.

## 🛠️ Go Concepts Demonstrated
* Standard library `time` package integration.
* Parsing and creating structured time values using `time.Date`.
* Calculating time difference spans via `time.Sub()`.
* Converting hours to days, and reading years and months from time structures.

## 📖 How to Run
Provide your birth year, month, and day as command-line arguments:
```bash
go run main.go 2000 12 15
```
Example Output:
```bash
Age in year: 25
Age in month: 306
Age in days: 9318
```