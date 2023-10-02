package repo

import (
	"database/sql"
	"log"
)

type UserDB struct {
	DB *sql.DB
}

func CreateUserRepository(db *sql.DB) *UserDB {
	return &UserDB{DB: db}
}

func AddUser(UserName string, Email string, hashedPassword string, db *sql.DB) {
	statement, err := db.Prepare("INSERT INTO users (Name, Email,Password) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	statement.Exec(UserName, Email, hashedPassword)
	db.Close()
}
