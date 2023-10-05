package controllers

import (
	"log"
	"net/http"
	"text/template"
)

func (controllers *Controllers) CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		controllers.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	if controllers.Service.UserService.CheckUserLogin(r) == false {
		http.Redirect(w, r, "/signin", 302)
	}

	cookie, err := r.Cookie("logged-in")
	if err != nil {
		controllers.ErrorHandler(w, http.StatusInternalServerError)
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
