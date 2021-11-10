package db

import (
	"fmt"
	"database/sql"
	"log"
	  _ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "mats"
	dbname   = "go-cli"
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

