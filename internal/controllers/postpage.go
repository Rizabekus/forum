package controllers

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

func PostPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	xurl := strings.Split(r.URL.String(), "id=")
	if len(xurl) < 2 {
		ErrorHandler(w, http.StatusNotFound)
		return
	}

	id, err := strconv.Atoi(xurl[1])
	if err != nil {
		ErrorHandler(w, http.StatusNotFound)
		return
	}
	db, err := sql.Open("sqlite3", "./sql/database.db")
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	qu, err := db.Query("select Title, Post,Namae from posts where Id=(?)", id)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	defer qu.Close()
	var title string
	var text string
	var name string
	for qu.Next() {
		qu.Scan(&title, &text, &name)
	}

	db.Close()
	if r.URL.String() != "/comments?id="+strconv.Itoa(id) {
		ErrorHandler(w, http.StatusNotFound)
		return
	}
	db, err = sql.Open("sqlite3", "./sql/database.db")
	defer db.Close()
	count, err := db.Query("select count(*) from posts;")
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	var i int
	defer count.Close()
	for count.Next() {
		count.Scan(&i)
	}
	if id > i || id < 1 {
		ErrorHandler(w, http.StatusNotFound)
		return
	}

	tmp, err := template.ParseFiles("./ui/html/comments.html")
	if err != nil {

		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	comments := internal.CollectComments(id, db)
	result := internal.Postpage{
		Title:    title,
		Post:     text,
		Name:     name,
		Comments: comments,
	}
	// fmt.Printf("%s i title\n%s is post\n%s is name\n", title, text, name)
	// fmt.Println(comments)
	err = tmp.Execute(w, result)
}
