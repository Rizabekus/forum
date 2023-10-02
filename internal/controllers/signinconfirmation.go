package controllers

import (
	"database/sql"
	"net/http"
	"text/template"
	"time"

	"github.com/gofrs/uuid"
)

func (controllers *Controllers) SignInConfirmation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		controllers.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	name := r.FormValue("UserName")
	password := r.FormValue("UserPassword")
	result, text := controllers.Service.UserService.ConfirmSignin(name, password)
	if result == true {
		u1, err := uuid.NewV4()
		if err != nil {
			controllers.ErrorHandler(w, http.StatusForbidden)
			return
		}
		u2 := uuid.NewV3(u1, name).String()
		db, err := sql.Open("sqlite3", "./sql/database.db")
		if err != nil {
			controllers.ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		defer db.Close()
		controllers.Service.UserService.CreateSession(u2, name, db)

		cookie := &http.Cookie{Name: "logged-in", Value: u2, Expires: time.Now().Add(365 * 24 * time.Hour)}
		http.SetCookie(w, cookie)
		http.Redirect(w, r, "/", 302)
	} else {
		tmpl, err := template.ParseFiles("./ui/html/signin.html")
		if err != nil {
			controllers.ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, text)
	}
}
