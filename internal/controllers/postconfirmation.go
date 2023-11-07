package controllers

import (
	"fmt"
	"forum/pkg"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
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

	// ErrMissingFile := errors.New("http: no such file")
	image, header, err := r.FormFile("image")
	if err != nil {
		if err == http.ErrMissingFile {
			ourImage, err := os.Open("./ui/NoImage.jpg")
			if err != nil {
				controllers.ErrorHandler(w, http.StatusInternalServerError)
				return
			}
			imageData, err := ioutil.ReadAll(ourImage)
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
				return
			} else {
				controllers.Service.PostService.CreatePost(cookie.Value, text, cat, title, imageData)
				http.Redirect(w, r, "/", 302)
				return
			}
		} else {
			controllers.ErrorHandler(w, http.StatusInternalServerError)
			return
		}
	}

	extention := strings.Split(header.Filename, ".")[len(strings.Split(header.Filename, "."))-1]

	if extention != "gif" && extention != "jpg" && extention != "png" && extention != "jpeg" {
		tmpl, err := template.ParseFiles("./ui/html/create.html")
		if err != nil {
			controllers.ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, "Wrong file for image! Only jpg, png, gif formats!")
		return
	}

	// Read the file content into a []byte

	imageData, err := ioutil.ReadAll(image)
	if err != nil {
		controllers.ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	image.Close()
	checker, t := pkg.PostChecker(title, text)
	if checker == false {
		tmpl, err := template.ParseFiles("./ui/html/create.html")
		if err != nil {
			controllers.ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, t)
		return
	} else {
		controllers.Service.PostService.CreatePost(cookie.Value, text, cat, title, imageData)
		http.Redirect(w, r, "/", 302)
		return
	}
}
