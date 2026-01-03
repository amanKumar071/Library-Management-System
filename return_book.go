package handlers

import (
	"GOLANG_PROJECT/database"
	"encoding/json"
	"net/http"
)

func ReturnBook(w http.ResponseWriter, r *http.Request) {
	var req struct {
		BookID int `json:"book_id"`
	}

	json.NewDecoder(r.Body).Decode(&req)

	database.DB.Exec(
		"UPDATE borrow_records SET status='returned' WHERE book_id=? AND status='borrowed'",
		req.BookID,
	)

	database.DB.Exec(
		"UPDATE books SET stock = stock + 1 WHERE id=?",
		req.BookID,
	)

	w.Write([]byte("Book returned successfully"))
}
