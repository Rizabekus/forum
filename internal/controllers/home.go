package controllers

import (
	"database/sql"
	"net/http"
	"text/template"
	"time"
)

func (controllers *Controllers) Homepage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		controllers.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		controllers.ErrorHandler(w, http.StatusNotFound)
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
			controllers.ErrorHandler(w, http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, controllers.Service.PostService.ShowPost())

	} else if err != nil {
		controllers.ErrorHandler(w, http.StatusInternalServerError)
		return
	} else {

		if time.Now().After(cookie.Expires) {
			// DeleteByID
			db, err := sql.Open("sqlite3", "./sql/database.db")

			tx, err := db.Begin()
			if err != nil {
				controllers.ErrorHandler(w, http.StatusInternalServerError)
				return
			}
			db.Exec("Delete * from cookies where Id = ( ? )", cookie.Value)
			tx.Commit()

			cookie = &http.Cookie{
				Name:  "logged-in",
				Value: "not-logged",
			}

		}

		c := cookie.Value
		db, err := sql.Open("sqlite3", "./sql/database.db")
		var name string
		// SelectUserByID
		Name, err := db.Query("SELECT lame FROM cookies WHERE Id = ( ? )", c)
		if err != nil {
			controllers.ErrorHandler(w, http.StatusInternalServerError)
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
			controllers.ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, controllers.Service.PostService.ShowPost())
	}
}
