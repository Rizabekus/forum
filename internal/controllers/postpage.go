package controllers

import (
	"forum/internal/models"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

func (controllers *Controllers) PostPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		controllers.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	xurl := strings.Split(r.URL.String(), "id=")
	if len(xurl) < 2 {
		controllers.ErrorHandler(w, http.StatusNotFound)
		return
	}

	id, err := strconv.Atoi(xurl[1])
	if err != nil {
		controllers.ErrorHandler(w, http.StatusNotFound)
		return
	}
	title, text, name := controllers.Service.PostService.SelectPostByID(id)

	if r.URL.String() != "/comments?id="+strconv.Itoa(id) {
		controllers.ErrorHandler(w, http.StatusNotFound)
		return
	}
	count := controllers.Service.PostService.CountPosts()

	if id > count || id < 1 {
		controllers.ErrorHandler(w, http.StatusNotFound)
		return
	}

	tmp, err := template.ParseFiles("./ui/html/comments.html")
	if err != nil {

		controllers.ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	comments := controllers.Service.CommentService.CollectComments(id)
	result := models.Postpage{
		Title:    title,
		Post:     text,
		Name:     name,
		Comments: comments,
	}
	// fmt.Printf("%s i title\n%s is post\n%s is name\n", title, text, name)
	// fmt.Println(comments)
	err = tmp.Execute(w, result)
}
