package controllers

import (
	"database/sql"
	"net/http"
	"text/template"
	"time"
)

func Homepage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		ErrorHandler(w, http.StatusNotFound)
		return
	}

	cookie, err := r.Cookie("logged-in")

	if err == http.ErrNoCookie || cookie.Value == "not-logged" {
		cookie = &http.Cookie{
			Name:  "logged-in",
			Value: "not-logged",
		}
		http.SetCookie(w, cookie)

		files := []string{
			"./ui/html/home.page.tmpl",
			"./ui/html/base.layout.tmpl",
		}

		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, repository.ShowPost())
	} else if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	} else {
		if time.Now().After(cookie.Expires) {
			db, err := sql.Open("sqlite3", "./sql/database.db")
			if err != nil {
				ErrorHandler(w, http.StatusInternalServerError)
				return
			}
			tx, err := db.Begin()
			if err != nil {
				ErrorHandler(w, http.StatusInternalServerError)
				return
			}
			db.Exec("Delete * from cookies where Id = ( ? )", cookie.Value)
			tx.Commit()
			db.Close()
			cookie = &http.Cookie{
				Name:  "logged-in",
				Value: "not-logged",
			}

		}
		c := cookie.Value
		db, err := sql.Open("sqlite3", "./sql/database.db")
		var name string

		Name, err := db.Query("SELECT lame FROM cookies WHERE Id = ( ? )", c)
		if err != nil {
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		defer Name.Close()
		for Name.Next() {
			Name.Scan(&name)
		}

		files := []string{
			"./ui/html/user.home.tmpl",
			"./ui/html/base.layout.tmpl",
		}
		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		db.Close()
		tmpl.Execute(w, repository.ShowPost())

	}
}
