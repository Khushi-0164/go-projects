# 🎮 Number Guessing Game
A simple interactive command-line guessing game. The computer generates a secret random number between 1 and 100, and evaluates whether your guess is too high, too low, or correct.

## 🚀 Features
* Random secret number generation.
* Single-guess CLI entry with bounds checking (1–100).
* Evaluation hints (`Too High` or `Too Low`).

## 🛠️ Go Concepts Demonstrated
* Random seeding using Unix nano timestamps (`rand.Seed(time.Now().UnixNano())`).
* Dynamic random range generation using `math/rand` (`rand.Intn(100)`).
* Parsing CLI input from `os.Args` to integer using `strconv.Atoi`.
* String formatting and interpolation with `fmt.Printf`.

## 📖 How to Run
Run the game by passing your guess as an argument:
```bash
go run main.go 42
```
Example Output:
```bash
Too Low : Secret was 76
```
