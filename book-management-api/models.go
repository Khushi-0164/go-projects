package main

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = []Book{
	{
		ID:     1,
		Title:  "Atomic Habits",
		Author: "James Clear",
	},
	{
		ID:     2,
		Title:  "The Alchemist",
		Author: "Paulo Coelho",
	},
}
