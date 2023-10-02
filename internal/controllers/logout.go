package controllers

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"
)

func (controllers *Controllers) Logout(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./sql/database.db")
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	cookie, err := r.Cookie("logged-in")
	internal.DeleteCookie(cookie.Value, db)
	cookie = &http.Cookie{
		Name:  "logged-in",
		Value: "not-logged",
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", 302)
}

func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	cookie, err := r.Cookie("logged-in")
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	if cookie.Value == "not-logged" {
		http.Redirect(w, r, "/signup", 302)
	} else {

		tmpl, err := template.ParseFiles("./ui/html/create.html")
		if err != nil {
			log.Println(err.Error())
			return
		}
		tmpl.Execute(w, nil)
	}
}
