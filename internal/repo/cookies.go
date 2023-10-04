package repo

import (
	"database/sql"
	"log"
)

type cookiesDB struct {
	DB *sql.DB
}

func CreateCookiesRepository(db *sql.DB) *cookiesDB {
	return &cookiesDB{DB: db}
}

func (db *cookiesDB) DeleteCookie(cookie string) {
	tx, err := db.DB.Begin()
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.DB.Exec("DELETE * FROM cookies WHERE Id=(?)", cookie)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
}
