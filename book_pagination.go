package handlers

import (
	"GOLANG_PROJECT/database"
	"GOLANG_PROJECT/models"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
)

func GetBooksPaginated(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	var rows *sql.Rows
	var err error

	baseQuery := `
	SELECT 
		b.id, b.title, b.isbn, b.publication_year, b.genre, b.stock,
		a.id, a.full_name
	FROM books b
	JOIN book_authors ba ON b.id = ba.book_id
	JOIN authors a ON ba.author_id = a.id
	ORDER BY b.id
	`

	
	if pageStr == "" && limitStr == "" {
		rows, err = database.DB.Query(baseQuery)
	} else {
		
		page, _ := strconv.Atoi(pageStr)
		limit, _ := strconv.Atoi(limitStr)

		if page < 1 {
			page = 1
		}
		if limit < 1 {
			limit = 5
		}

		offset := (page - 1) * limit

		paginatedQuery := baseQuery + " LIMIT ? OFFSET ?"
		rows, err = database.DB.Query(paginatedQuery, limit, offset)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	bookMap := make(map[int]*models.Book)

	for rows.Next() {
		var bookID, authorID, stock int
		var title, isbn, genre, authorName string
		var year int

		rows.Scan(
			&bookID, &title, &isbn, &year, &genre, &stock,
			&authorID, &authorName,
		)

		if _, exists := bookMap[bookID]; !exists {
			bookMap[bookID] = &models.Book{
				ID:              bookID,
				Title:           title,
				ISBN:            isbn,
				PublicationYear: year,
				Genre:           genre,
				Stock:           stock,
				Authors:         []models.Author{},
			}
		}

		bookMap[bookID].Authors = append(
			bookMap[bookID].Authors,
			models.Author{
				ID:   authorID,
				Name: authorName,
			},
		)
	}

	books := []models.Book{}
	for _, b := range bookMap {
		books = append(books, *b)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
