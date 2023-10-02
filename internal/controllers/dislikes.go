package controllers

import (
	"database/sql"
	"net/http"
)

func (controllers *Controllers) Dislikes(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		controllers.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	id := r.FormValue("id")

	db, err := sql.Open("sqlite3", "./sql/database.db")
	if err != nil {
		controllers.ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	cookie, err := r.Cookie("logged-in")
	if err != nil {
		controllers.ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	var checkName string
	row, err := db.Query("SELECT lame FROM cookies WHERE Id=(?)", cookie.Value)
	for row.Next() {
		row.Scan(&checkName)
	}
	row.Close()
	db, err = sql.Open("sqlite3", "./sql/database.db")
	if err != nil {
		controllers.ErrorHandler(w, http.StatusInternalServerError)
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
	if checklikes == true && checkdislikes == false {

		tx, err := db.Begin()
		if err != nil {
			controllers.ErrorHandler(w, http.StatusInternalServerError)
			return
		}

		_, err = db.Exec("INSERT INTO dislikes (Name, Postid) VALUES (?, ?)", checkName, id)
		_, err = db.Exec("DELETE FROM likes WHERE Name=(?) and Postid = (?)", checkName, id)
		if err != nil {
			controllers.ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		tx.Commit()
		db.Close()
	} else if checklikes == false && checkdislikes == false {
		tx, err := db.Begin()
		if err != nil {
			controllers.ErrorHandler(w, http.StatusInternalServerError)
			return
		}

		_, err = db.Exec("INSERT INTO dislikes (Name, Postid) VALUES (?, ?)", checkName, id)

		if err != nil {
			controllers.ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		tx.Commit()
		db.Close()
	} else if checklikes == false && checkdislikes == true {
		tx, err := db.Begin()
		if err != nil {
			controllers.ErrorHandler(w, http.StatusInternalServerError)
			return
		}

		_, err = db.Exec("DELETE FROM dislikes WHERE Name=(?) and Postid=(?)", checkName, id)

		if err != nil {
			controllers.ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		tx.Commit()
		db.Close()

	}

	http.Redirect(w, r, previousURL, 302)
}
