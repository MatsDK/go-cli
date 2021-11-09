package db

import (
	"fmt"
	"database/sql"

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
	//defer db.Close()

	return db
}

