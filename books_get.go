package handlers

import (
	"GOLANG_PROJECT/database"
	"GOLANG_PROJECT/models"
	"encoding/json"
	"net/http"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	query := `
	SELECT 
		b.id, b.title, b.isbn, b.publication_year, b.genre,
		a.id, a.full_name
	FROM books b
	JOIN book_authors ba ON b.id = ba.book_id
	JOIN authors a ON ba.author_id = a.id
	ORDER BY b.id;
	`

	rows, err := database.DB.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	bookMap := make(map[int]*models.Book)

	for rows.Next() {
		var bookID, authorID int
		var title, isbn, genre, authorName string
		var year int

		rows.Scan(&bookID, &title, &isbn, &year, &genre, &authorID, &authorName)

		if _, exists := bookMap[bookID]; !exists {
			bookMap[bookID] = &models.Book{
				ID:              bookID,
				Title:           title,
				ISBN:            isbn,
				PublicationYear: year,
				Genre:           genre,
				Authors:         []models.Author{},
			}
		}

		bookMap[bookID].Authors = append(
			bookMap[bookID].Authors,
			models.Author{ID: authorID, Name: authorName},
		)
	}

	books := []models.Book{}
	for _, b := range bookMap {
		books = append(books, *b)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
