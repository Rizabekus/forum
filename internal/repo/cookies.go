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
	_, err := db.DB.Exec("DELETE FROM cookies WHERE Id=(?)", cookie)
	if err != nil {
		log.Fatal(err)
	}
}
