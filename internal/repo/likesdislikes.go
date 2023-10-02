package repo

import "database/sql"

type likesdislikesDB struct {
	DB *sql.DB
}

func CreateLikesdislikesRepository(db *sql.DB) *likesdislikesDB {
	return &likesdislikesDB{DB: db}
}
