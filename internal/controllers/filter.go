package controllers

import (
	"net/http"
	"text/template"
)

func (controllers *Controllers) Filter(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {

		w.Header().Set("Allow", http.MethodPost)
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
		return
	}
	r.ParseForm()
	cc := controllers.Service.CookiesService.GetCookie(r)

	namecookie := controllers.Service.UserService.FindUserByToken(cc.Value)

	likesdislikes := r.Form["LikeDislike"]
	categories := r.Form["Category"]
	yourposts := r.Form["YourPosts"]
	if len(likesdislikes) == 0 && len(categories) == 0 && len(yourposts) == 0 {
		http.Redirect(w, r, "/", 301)
	}
	posts := controllers.Service.PostService.Filter(namecookie, likesdislikes, categories, yourposts)

	cook, err := r.Cookie("logged-in")
	var files []string
	if err == http.ErrNoCookie || cook.Value == "not-logged" {
		files = []string{
			"./ui/html/home.page.tmpl",
			"./ui/html/base.layout.tmpl",
		}
	} else {
		files = []string{
			"./ui/html/user.home.tmpl",
			"./ui/html/base.layout.tmpl",
		}
	}
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		controllers.ErrorHandler(w, http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, posts)
}
