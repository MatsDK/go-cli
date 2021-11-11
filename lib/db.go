package lib

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     string = "localhost"
	port     int    = 5432
	user     string = "postgres"
	password string = "mats"
	dbname   string = "go-cli"
)

func ConnectDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	return db
}

func Query(query string, db *sql.DB) *sql.Rows {
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("SQL select error: ")
		log.Fatal(err)
	}

	return rows
}
