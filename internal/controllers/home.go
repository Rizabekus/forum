package controllers

import (
	"net/http"
	"text/template"
)

func (controllers *Controllers) Homepage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		controllers.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		controllers.ErrorHandler(w, http.StatusNotFound)
		return
	}

	cookie, err := r.Cookie("logged-in")

	if err == http.ErrNoCookie || cookie.Value == "not-logged" {

		cookie = &http.Cookie{
			Name:  "logged-in",
			Value: "not-logged",
		}
		http.SetCookie(w, cookie)

		files := []string{
			"./ui/html/home.page.tmpl",
			"./ui/html/base.layout.tmpl",
		}

		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			controllers.ErrorHandler(w, http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, controllers.Service.PostService.ShowPost())

	} else if err != nil {
		controllers.ErrorHandler(w, http.StatusInternalServerError)
		return
	} else {

		// if time.Now().After(cookie.Expires) {
		// 	fmt.Println("GGG")
		// 	controllers.Service.CookiesService.DeleteCookie(cookie.Value)

		// 	cookie = &http.Cookie{
		// 		Name:  "logged-in",
		// 		Value: "not-logged",
		// 	}

		// }

		files := []string{
			"./ui/html/user.home.tmpl",
			"./ui/html/base.layout.tmpl",
		}
		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			controllers.ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, controllers.Service.PostService.ShowPost())
	}
}
