package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Expense struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Amount float64 `json:"amount"`
	Date   string  `json:"date"`
}

func GetNextID(expenses []Expense) int {
	maxID := 0

	for _, expense := range expenses {
		if expense.ID > maxID {
			maxID = expense.ID
		}
	}

	return maxID + 1
}

func LoadExpenses() ([]Expense, error) {
	data, err := os.ReadFile("expenses.json")
	if err != nil {
		return []Expense{}, nil
	}

	var expenses []Expense

	err = json.Unmarshal(data, &expenses)
	if err != nil {
		return nil, err
	}

	return expenses, nil
}

func SaveExpenses(expenses []Expense) error {
	data, err := json.MarshalIndent(expenses, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile("expenses.json", data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func ListExpenses(expenses []Expense) {
	if len(expenses) == 0 {
		fmt.Println("No expenses found")
		return
	}

	for _, expense := range expenses {
		fmt.Printf(
			"ID: %d | Title: %s | Amount: %.2f | Date: %s\n",
			expense.ID,
			expense.Title,
			expense.Amount,
			expense.Date,
		)
	}
}

func AddExpenses(expenses []Expense, title string, amount float64) []Expense {
	newExpense := Expense{
		ID:     GetNextID(expenses),
		Title:  title,
		Amount: amount,
		Date:   time.Now().Format("2006-01-02"),
	}

	expenses = append(expenses, newExpense)

	fmt.Println("Expense added successfully")

	return expenses
}

func DeleteExpense(expenses []Expense, id int) []Expense {
	for i, expense := range expenses {
		if expense.ID == id {
			expenses = append(expenses[:i], expenses[i+1:]...)
			fmt.Println("Expense deleted")
			return expenses
		}
	}

	fmt.Println("Expense not found")
	return expenses
}

func SummaryExpenses(expenses []Expense) {
	total := 0.0

	for _, expense := range expenses {
		total += expense.Amount
	}

	fmt.Printf("Total Expenses: %.2f\n", total)
	fmt.Printf("Number of Expenses: %d\n", len(expenses))
}

func main() {
	expenses, err := LoadExpenses()
	if err != nil {
		fmt.Println("Error loading expenses:", err)
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("go run main.go add \"title\" amount")
		fmt.Println("go run main.go list")
		fmt.Println("go run main.go delete <id>")
		fmt.Println("go run main.go summary")
		return
	}

	command := os.Args[1]

	switch command {

	case "list":
		ListExpenses(expenses)
		return

	case "add":
		if len(os.Args) < 4 {
			fmt.Println("Usage: go run main.go add \"title\" amount")
			return
		}

		title := os.Args[2]

		amount, err := strconv.ParseFloat(os.Args[3], 64)
		if err != nil {
			fmt.Println("Invalid amount")
			return
		}

		expenses = AddExpenses(expenses, title, amount)

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run main.go delete <id>")
			return
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid ID")
			return
		}

		expenses = DeleteExpense(expenses, id)

	case "summary":
		SummaryExpenses(expenses)
		return

	default:
		fmt.Println("Invalid command")
		return
	}

	err = SaveExpenses(expenses)
	if err != nil {
		fmt.Println("Error saving expenses:", err)
	}
}
