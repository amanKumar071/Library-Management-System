package main

import (
	"log"
	"net/http"

	"GOLANG_PROJECT/database"
	"GOLANG_PROJECT/handlers"
)

func main() {
	database.ConnectDB()

	
	http.HandleFunc("/books", handlers.GetBooksPaginated)

	http.HandleFunc("/books/add", handlers.AddBook)
	http.HandleFunc("/borrowers/add", handlers.AddBorrower)
	http.HandleFunc("/borrow", handlers.BorrowBook)
	http.HandleFunc("/return", handlers.ReturnBook)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
