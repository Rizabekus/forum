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

func (UserDB *UserDB) AddUser(UserName string, Email string, hashedPassword string) {
	statement, err := UserDB.DB.Prepare("INSERT INTO users (Name, Email,Password) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	statement.Exec(UserName, Email, hashedPassword)
	UserDB.DB.Close()
}

func (UserDB *UserDB) CreateSession(id, name string) {
	tx, err := UserDB.DB.Begin()
	if err != nil {
		log.Fatal(err)
	}

	_, err = UserDB.DB.Exec("INSERT INTO cookies (Id, lame) VALUES (?, ?)", id, name)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
	UserDB.DB.Close()
}
