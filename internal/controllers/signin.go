package controllers

import (
	"net/http"
	"text/template"
)

func (controllers *Controllers) SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	tmpl, err := template.ParseFiles("./ui/html/signin.html")
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
