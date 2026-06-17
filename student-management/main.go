package main

import (
	"fmt"
	"os"
	"strconv"

	"student-management/service"
	"student-management/storage"
)

func main() {

	store := storage.JSONStorage{
		FileName: "students.json",
	}

	studentService := service.StudentService{
		Storage: store,
	}

	if len(os.Args) < 2 {
		fmt.Println("Please provide a command")
		return
	}

	command := os.Args[1]

	switch command {

	case "add":

		if len(os.Args) != 4 {
			fmt.Println("Usage: go run . add <name> <age>")
			return
		}

		name := os.Args[2]

		age, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println("Age must be a number")
			return
		}

		err = studentService.AddStudent(name, age)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Student added successfully")

	case "list":

		students, err := studentService.ListStudents()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		for _, student := range students {
			fmt.Printf("ID: %d | Name: %s | Age: %d\n",
				student.ID,
				student.Name,
				student.Age,
			)
		}

	case "delete":
		if len(os.Args) != 3 {
			fmt.Println("Usage: go run . delete <id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Id must be a number")
			return
		}
		err = studentService.DeleteStudent(id)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Student Deleted Successfully")

	case "update":
		if len(os.Args) != 5 {
			fmt.Println("Usage: go run . update <id> <name> <age>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Id must be a number")
			return
		}
		name := os.Args[3]

		age, err := strconv.Atoi(os.Args[4])
		if err != nil {
			fmt.Println("Age must be a number")
			return
		}
		err = studentService.UpdateStudent(id, name, age)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Student Updated Successfully")

	default:
		fmt.Println("Unknown command")
	}
}
