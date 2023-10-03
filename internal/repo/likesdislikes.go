package repo

import (
	"database/sql"
	"log"
)

type likesdislikesDB struct {
	DB *sql.DB
}

func CreateLikesDislikesRepository(db *sql.DB) *likesdislikesDB {
	return &likesdislikesDB{DB: db}
}

func (db *likesdislikesDB) CheckLikeExistence(user string, id string) bool {
	checklikes := false
	rows, err := db.DB.Query("SELECT Name FROM likes WHERE Postid=(?)", id)
	if err != nil {
		log.Fatal(err)
	}
	var likerName string
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&likerName)
		if likerName == user {
			checklikes = true
		}
	}
	return checklikes
}

func (db *likesdislikesDB) CheckDislikeExistence(user string, id string) bool {
	checkdislikes := false
	x, err := db.DB.Query("SELECT Name FROM dislikes WHERE Postid=(?)", id)
	if err != nil {
		log.Fatal(err)
	}
	var dislikerName string
	defer x.Close()

	for x.Next() {
		x.Scan(&dislikerName)
		if dislikerName == user {
			checkdislikes = true
		}
	}
	return checkdislikes
}

func (db *likesdislikesDB) AddLike(user string, id string) {
	tx, err := db.DB.Begin()
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.DB.Exec("INSERT INTO likes (Name, Postid) VALUES (?, ?)", user, id)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
}

func (db *likesdislikesDB) AddDislike(user string, id string) {
	tx, err := db.DB.Begin()
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.DB.Exec("INSERT INTO dislikes (Name, Postid) VALUES (?, ?)", user, id)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
}

func (db *likesdislikesDB) RemoveLike(user string, id string) {
	tx, err := db.DB.Begin()
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.DB.Exec("DELETE FROM likes WHERE Name=(?) and Postid=(?)", user, id)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
}

func (db *likesdislikesDB) RemoveDislike(user string, id string) {
	tx, err := db.DB.Begin()
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.DB.Exec("DELETE FROM dislikes WHERE Name=(?) and Postid=(?)", user, id)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
}
