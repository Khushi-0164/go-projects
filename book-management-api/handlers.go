package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Book Management API!")
}

func booksHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		id := r.URL.Query().Get("id")

		if id == "" {
			getBooks(w, r)
		} else {
			getBookByID(w, r)
		}

	case http.MethodPost:
		addBook(w, r)

	case http.MethodDelete:
		deleteBook(w, r)

	case http.MethodPut:
		updateBook(w, r)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, title, author FROM books")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var books []Book

	for rows.Next() {
		var book Book

		err := rows.Scan(&book.ID, &book.Title, &book.Author)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		books = append(books, book)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func addBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newBook Book

	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err = db.QueryRow(
		"INSERT INTO books (title,author) VALUES ($1,$2) RETURNING id",
		newBook.Title,
		newBook.Author,
	).Scan(&newBook.ID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBook)
}

func getBookByID(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var book Book
	err = db.QueryRow(
		"SELECT id,title,author FROM books WHERE id=$1",
		id,
	).Scan(&book.ID, &book.Title, &book.Author)

	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	result, err := db.Exec(
		"DELETE FROM books WHERE id=$1",
		id,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func updateBook(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updatedBook Book

	err = json.NewDecoder(r.Body).Decode(&updatedBook)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	result, err := db.Exec(
		"UPDATE books SET title=$1,author=$2 WHERE id=$3",
		updatedBook.Title,
		updatedBook.Author,
		id,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowAffected, _ := result.RowsAffected()

	if rowAffected == 0 {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	updatedBook.ID = id

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedBook)

}
