package handlers

import (
	"GOLANG_PROJECT/database"
	"encoding/json"
	"net/http"
)

func BorrowBook(w http.ResponseWriter, r *http.Request) {
	var req struct {
		BookID     int `json:"book_id"`
		BorrowerID int `json:"borrower_id"`
	}

	json.NewDecoder(r.Body).Decode(&req)

	var stock int
	database.DB.QueryRow(
		"SELECT stock FROM books WHERE id = ?",
		req.BookID,
	).Scan(&stock)

	if stock <= 0 {
		http.Error(w, "Book out of stock", http.StatusBadRequest)
		return
	}

	database.DB.Exec(
		"INSERT INTO borrow_records (book_id, borrower_id, status) VALUES (?, ?, 'borrowed')",
		req.BookID,
		req.BorrowerID,
	)

	database.DB.Exec(
		"UPDATE books SET stock = stock - 1 WHERE id = ?",
		req.BookID,
	)

	w.Write([]byte("Book borrowed successfully"))
}
