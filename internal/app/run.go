package app

import (
	"database/sql"
	"log"
)

func Run() {
	db, err := sql.Open("sqlite3", "databse.db")
	if err != nil {
		log.Fatal(err)
	}
	Server()
	defer db.Close()
}

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}
