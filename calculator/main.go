package main

import (
	"fmt"
	"os"
	"strconv"
)

func Addition(a, b int) int {
	return a + b

}
func Subtraction(a, b int) int {
	return a - b
}
func Multiply(a, b int) int {
	return a * b
}
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Cannot divide by 0")
	}
	return a / b, nil
}
func main() {
	if len(os.Args) < 4 {
		fmt.Println("Plese Provide correct input")
		return
	}

	num1, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("First Number must be integer")
		return
	}
	num2, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("Second Number must be integer")
		return
	}

	switch os.Args[1] {
	case "add":
		fmt.Println("Output:", Addition(num1, num2))
	case "sub":
		fmt.Println("Output:", Subtraction(num1, num2))
	case "mul":
		fmt.Println("Output:", Multiply(num1, num2))
	case "divide":
		result, err := Divide(num1, num2)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Output:", result)
	default:
		fmt.Println("Invalid Command")
		fmt.Println("Available commands: add, sub, mul, divide")
	}
}
