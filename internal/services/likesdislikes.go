package services

import (
	"database/sql"
	"forum/internal/models"
	"net/http"
)

type LikesDislikesService struct {
	repo models.LikesDislikesRepository
}

func CreateLikesDislikesService(repo models.LikesDislikesRepository) *LikesDislikesService {
	return &LikesDislikesService{repo: repo}
}

func (LikesDislikesService *LikesDislikesService) Like(id int, cookie string) {
	var checkName string
	row, err := db.Query("SELECT lame FROM cookies WHERE Id=(?)", cookie.Value)
	for row.Next() {
		row.Scan(&checkName)
	}
	row.Close()
	db, err = sql.Open("sqlite3", "./sql/database.db")
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}

	previousURL := r.Header.Get("Referer")
	checklikes := false
	checkdislikes := false
	rows, err := db.Query("SELECT Name FROM likes WHERE Postid=(?)", id)
	var likerName string
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&likerName)
		if likerName == checkName {
			checklikes = true
		}
	}
	x, err := db.Query("SELECT Name FROM dislikes WHERE Postid=(?)", id)
	var dislikerName string
	defer x.Close()

	for x.Next() {
		x.Scan(&dislikerName)
		if dislikerName == checkName {
			checkdislikes = true
		}
	}
	if checklikes == false && checkdislikes == true {

		tx, err := db.Begin()
		if err != nil {
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}

		_, err = db.Exec("INSERT INTO likes (Name, Postid) VALUES (?, ?)", checkName, id)
		_, err = db.Exec("DELETE FROM dislikes WHERE Name=(?) and Postid=(?)", checkName, id)
		if err != nil {
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		tx.Commit()
		db.Close()
	} else if checklikes == false && checkdislikes == false {
		tx, err := db.Begin()
		if err != nil {
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}

		_, err = db.Exec("INSERT INTO likes (Name, Postid) VALUES (?, ?)", checkName, id)

		if err != nil {
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		tx.Commit()
		db.Close()
	} else if checklikes == true && checkdislikes == false {
		tx, err := db.Begin()
		if err != nil {
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}

		_, err = db.Exec("DELETE FROM likes WHERE Name=(?) and Postid=(?)", checkName, id)

		if err != nil {
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		tx.Commit()
		db.Close()
	}
}
