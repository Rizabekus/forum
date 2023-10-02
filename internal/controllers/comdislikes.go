package controllers

import (
	"database/sql"
	"net/http"
	"strings"
)

func (controllers *Controllers) ComDislikes(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		controllers.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	previousURL := r.Header.Get("Referer")
	postid := (strings.Split(previousURL, "id="))[1]
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
	checklikes := false
	checkdislikes := false
	rows, err := db.Query("SELECT Name FROM comlikes WHERE (Comid,Id)=(?,?)", id, postid)
	if err != nil {
		controllers.ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	var likerName string
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&likerName)
		if likerName == checkName {
			checklikes = true
		}
	}
	x, err := db.Query("SELECT Name FROM comdislikes WHERE (Comid,Id)=(?,?)", id, postid)
	if err != nil {
		controllers.ErrorHandler(w, http.StatusInternalServerError)
		return
	}

	var dislikerName string
	defer x.Close()

	for x.Next() {
		x.Scan(&dislikerName)
		if dislikerName == checkName {
			checkdislikes = true
		}
	}
	tx, err := db.Begin()
	if err != nil {
		controllers.ErrorHandler(w, http.StatusInternalServerError)
		return
	}

	if checklikes == true && checkdislikes == false {

		_, err = db.Exec("INSERT INTO comdislikes (Name, Comid,Id) VALUES (?, ?, ?)", checkName, id, postid)
		_, err = db.Exec("DELETE FROM comlikes WHERE Name=(?) and Comid=(?) and Id=(?)", checkName, id, postid)
		if err != nil {
			controllers.ErrorHandler(w, http.StatusInternalServerError)
			return
		}

	} else if checklikes == false && checkdislikes == false {

		_, err = db.Exec("INSERT INTO comdislikes (Name, Comid,Id) VALUES (?, ?, ?)", checkName, id, postid)

		if err != nil {
			controllers.ErrorHandler(w, http.StatusInternalServerError)
			return
		}

	} else if checklikes == false && checkdislikes == true {

		_, err = db.Exec("DELETE FROM comdislikes WHERE Name=(?) and Comid=(?) and Id=(?)", checkName, id, postid)

		if err != nil {
			controllers.ErrorHandler(w, http.StatusInternalServerError)
			return
		}
	}
	tx.Commit()

	http.Redirect(w, r, previousURL, 302)
}
