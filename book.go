package models

type Author struct {
	ID   int    `json:"id"`
	Name string `json:"full_name"`
}

type Book struct {
	ID              int      `json:"id"`
	Title           string   `json:"title"`
	ISBN            string   `json:"isbn"`
	PublicationYear int      `json:"publication_year"`
	Genre           string   `json:"genre"`
	Authors         []Author `json:"authors"`
	Stock           int      `json:"stock"`
}

type Borrower struct {
	ID       int    `json:"id"` // primary key
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"` // corresponds to phone_number column
}
