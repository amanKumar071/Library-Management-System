package handlers

import (
	"GOLANG_PROJECT/database"
	"GOLANG_PROJECT/models"
	"encoding/json"
	"net/http"
)

func AddBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	query := `
	INSERT INTO books (title, isbn, publication_year, genre, stock)
	VALUES (?, ?, ?, ?, ?)
	`

	result, err := database.DB.Exec(
		query,
		book.Title,
		book.ISBN,
		book.PublicationYear,
		book.Genre,
		book.Stock,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	book.ID = int(id)

	json.NewEncoder(w).Encode(book)
}
