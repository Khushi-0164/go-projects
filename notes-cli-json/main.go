package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Note struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func GetNextID(notes []Note) int {
	maxID := 0

	for _, note := range notes {
		if note.ID > maxID {
			maxID = note.ID
		}
	}

	return maxID + 1
}

func LoadNotes() ([]Note, error) {
	data, err := os.ReadFile("notes.json")
	if err != nil {
		return []Note{}, nil
	}
	var notes []Note
	err = json.Unmarshal(data, &notes)
	if err != nil {
		return nil, err
	}
	return notes, err
}
func SaveNotes(notes []Note) error {
	data, err := json.MarshalIndent(notes, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile("notes.json", data, 0644)
	if err != nil {
		return err
	}
	return err
}
func AddNotes(notes []Note, title string, body string) []Note {
	newNotes := Note{
		ID:    GetNextID(notes),
		Title: title,
		Body:  body,
	}
	notes = append(notes, newNotes)
	return notes
}
func ListNotes(notes []Note) {
	if len(notes) == 0 {
		fmt.Println("No notes found")
		return
	}
	for _, note := range notes {
		fmt.Println(note.ID)
		fmt.Println(note.Title)
		fmt.Println(note.Body)
	}

}
func DeleteNotes(notes []Note, id int) []Note {
	if len(notes) == 0 {
		fmt.Println("No notes found")
		return []Note{}
	}
	for i, note := range notes {
		if note.ID == id {
			notes = append(notes[:i], notes[i+1:]...)
			fmt.Println("Notes deleted")
			return notes
		}
	}

	fmt.Println("Notes not found")
	return notes

}

func UpdateNotes(notes []Note, id int, title string, body string) []Note {
	if len(notes) == 0 {
		fmt.Println("No notes found")
		return []Note{}
	}
	for i, note := range notes {
		if note.ID == id {
			notes[i].Title = title
			notes[i].Body = body
			fmt.Println("Notes updated")
			return notes
		}
	}

	fmt.Println("Notes not found")
	return notes

}

func SearchNotes(notes []Note, keyword string) []Note {
	if len(notes) == 0 {
		fmt.Println("No notes found")
		return []Note{}
	}
	for _, note := range notes {
		if strings.Contains(note.Title, keyword) ||
			strings.Contains(note.Body, keyword) {

			fmt.Println(note.ID, note.Title, note.Body)
		}
	}
	return notes
}

func main() {
	notes, err := LoadNotes()
	if err != nil {
		fmt.Println("Error loading notes:", err)
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("go run main.go add \"title\" body")
		fmt.Println("go run main.go list")
		fmt.Println("go run main.go delete <id>")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 4 {
			fmt.Println("Usage: go run main.go add \"title\" body")
			return
		}
		title := os.Args[2]
		body := os.Args[3]
		notes = AddNotes(notes, title, body)

	case "list":
		ListNotes(notes)
		return

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run main.go delete <id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid Id")
			return
		}
		notes = DeleteNotes(notes, id)

	case "update":
		if len(os.Args) < 5 {
			fmt.Println("Usage: go run main.go update <id><title><body>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid Id")
			return
		}
		title := os.Args[3]
		body := os.Args[4]
		notes = UpdateNotes(notes, id, title, body)

	case "search":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run main.go search <title>")
			return
		}
		keyword := os.Args[2]
		notes = SearchNotes(notes, keyword)

	default:
		fmt.Println("Invalid command")
		return
	}

	err = SaveNotes(notes)
	if err != nil {
		fmt.Println("Error saving notes:", err)
		return
	}
}
