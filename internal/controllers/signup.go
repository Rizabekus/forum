package controllers

import (
	"log"
	"net/http"
	"text/template"
)

func (controllers *Controllers) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		controllers.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	tmpl, err := template.ParseFiles("./ui/html/signup.html")
	if err != nil {
		log.Println(err.Error())
		return
	}
	tmpl.Execute(w, nil)
}
