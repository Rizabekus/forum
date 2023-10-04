package controllers

import (
	"fmt"
	"net/http"
	"text/template"

	"golang.org/x/crypto/bcrypt"
)

func (controllers *Controllers) SignUpConfirmation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		controllers.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}

	name := r.FormValue("UserName")

	email := r.FormValue("UserEmail")
	password := r.FormValue("UserPassword")
	rewrittenPassword := r.FormValue("UserRewrittenPassword")
	fmt.Println("works")
	result, text := controllers.Service.UserService.ConfirmSignup(name, email, password, rewrittenPassword)

	if result == true {

		pwd, err := bcrypt.GenerateFromPassword([]byte(password), 1)
		if err != nil {
			controllers.ErrorHandler(w, http.StatusInternalServerError)
			return
		}

		controllers.Service.UserService.AddUser(name, email, string(pwd))

		http.Redirect(w, r, "/signin", 302)
	} else {
		tmpl, err := template.ParseFiles("./ui/html/signup.html")
		if err != nil {
			controllers.ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, text)
	}
}
