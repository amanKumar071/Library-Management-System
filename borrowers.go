package handlers

import (
	"GOLANG_PROJECT/database"
	"GOLANG_PROJECT/models"
	"encoding/json"
	"net/http"
)

func AddBorrower(w http.ResponseWriter, r *http.Request) {
	var b models.Borrower

	json.NewDecoder(r.Body).Decode(&b)

	query := `
	INSERT INTO borrowers (full_name, email, phone_number)
	VALUES (?, ?, ?)
	`

	_, err := database.DB.Exec(query, b.FullName, b.Email, b.Phone)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(b)
}
