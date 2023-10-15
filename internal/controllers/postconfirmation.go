package controllers

import (
	"fmt"
	"forum/pkg"
	"net/http"
	"text/template"
)

func (controllers *Controllers) PostConfirmation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		controllers.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	// cookie, err := r.Cookie("logged-in")

	title := r.FormValue("title")
	text := r.FormValue("convert")
	// cat := r.FormValue("cars")
	image, _, err := r.FormFile("image")
	if err != nil {
		controllers.ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	fmt.Println(image)
	checker, t := pkg.PostChecker(title, text)
	if checker == false {
		tmpl, err := template.ParseFiles("./ui/html/create.html")
		if err != nil {
			controllers.ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, t)
	} else {
		// controllers.Service.PostService.CreatePost(cookie.Value, text, cat, title, image)
		http.Redirect(w, r, "/", 302)
	}
}
