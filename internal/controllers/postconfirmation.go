package controllers

import (
	"fmt"
	"forum/pkg"
	"io/ioutil"
	"net/http"
	"text/template"
)

func (controllers *Controllers) PostConfirmation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		controllers.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}

	cookie, err := r.Cookie("logged-in")
	if err != nil {
		controllers.ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	err = r.ParseMultipartForm(20 << 20) // max size 20MB
	if err != nil {
		// over 20mb
		fmt.Println(err)
		controllers.ErrorHandler(w, http.StatusInternalServerError)
		return
	}

	err = r.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}
	title := r.FormValue("title")
	text := r.FormValue("convert")
	cat := r.FormValue("cars")

	image, _, err := r.FormFile("image")
	if err != nil {
		fmt.Println(err)
		controllers.ErrorHandler(w, http.StatusInternalServerError)
		return
	}

	defer image.Close()

	// Read the file content into a []byte

	imageData, err := ioutil.ReadAll(image)
	if err != nil {
		controllers.ErrorHandler(w, http.StatusInternalServerError)
		return
	}

	checker, t := pkg.PostChecker(title, text)
	if checker == false {
		tmpl, err := template.ParseFiles("./ui/html/create.html")
		if err != nil {
			controllers.ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, t)
	} else {
		controllers.Service.PostService.CreatePost(cookie.Value, text, cat, title, imageData)
		http.Redirect(w, r, "/", 302)
	}
}
