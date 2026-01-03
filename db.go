package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	var err error

	//  database name is `library`
	DB, err = sql.Open(
		"mysql",
		"root:1234@tcp(127.0.0.1:3306)/library",
	)

	if err != nil {
		log.Fatal(err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MySQL database: library")
}
