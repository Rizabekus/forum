package controllers

import (
	"net/http"
	"text/template"
)

func (controllers *Controllers) PostConfirmation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	cookie, err := r.Cookie("logged-in")
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}

	title := r.FormValue("title")
	text := r.FormValue("convert")
	cat := r.FormValue("cars")
	checker, t := internal.PostChecker(title, text)
	if checker == false {
		tmpl, err := template.ParseFiles("./ui/html/create.html")
		if err != nil {
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, t)
	} else {

		internal.CreatePost(cookie.Value, text, cat, title)
		http.Redirect(w, r, "/", 302)
	}
}
