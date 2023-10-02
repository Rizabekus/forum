package controllers

import (
	"database/sql"
	"fmt"
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
	cookie, err := r.Cookie("logged-in")
	cc := cookie.Value
	db, err := sql.Open("sqlite3", "./sql/database.db")
	var namecookie string

	Name, err := db.Query("SELECT lame FROM cookies WHERE Id = ( ? )", cc)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	defer Name.Close()
	for Name.Next() {
		Name.Scan(&namecookie)
	}
	Name.Close()
	db.Close()

	likesdislikes := r.Form["LikeDislike"]
	categories := r.Form["Category"]
	var formattedlikes []string

	for i := range likesdislikes {
		formattedlikes = append(formattedlikes, likesdislikes[i]+"s.Postid")
	}
	// all := append(likes, categories...)
	if len(likesdislikes) == 0 && len(categories) == 0 && len(r.Form["YourPosts"]) == 0 {
		http.Redirect(w, r, "/", 301)
	}

	text := "SELECT posts.Title, posts.Post, posts.Namae, posts.Category, posts.Id from posts "

	if len(likesdislikes) == 2 {
		text = text + "INNER JOIN likes on posts.Id=likes.Postid INNER JOIN dislikes on posts.Id=dislikes.Postid"
	} else if len(likesdislikes) == 1 {
		if likesdislikes[0] == "like" {
			text = text + "INNER JOIN likes on posts.Id=likes.Postid"
		} else {
			text = text + "INNER JOIN dislikes on posts.Id=dislikes.Postid"
		}
	}
	if len(likesdislikes) > 0 {
		// text = text + " WHERE posts.Id IN (" + strings.Join(formattedlikes, ", ") + ")"
		text = text + " WHERE "
		for i := range formattedlikes {
			if i == 0 {
				text = text + "(posts.Id=" + formattedlikes[i] + " AND " + likesdislikes[i] + "s.Name=\"" + namecookie + "\")"
			} else {
				text = text + " OR (posts.Id=" + formattedlikes[i] + " AND " + likesdislikes[i] + "s.Name=\"" + namecookie + "\")"
			}
		}
	} else if len(categories) > 0 {
		text = text + " WHERE "
		for i := range categories {
			// text = text + " OR posts.Category=\"" + categories[i] + "\""
			if i == 0 {
				text = text + "posts.Category=\"" + categories[i] + "\""
			} else {
				text = text + " OR posts.Category=\"" + categories[i] + "\""
			}
		}
	}
	if len(categories) > 0 {
		if len(likesdislikes) > 0 {
			text = text + " AND ("
			for i := range categories {
				// text = text + " OR posts.Category=\"" + categories[i] + "\""
				if i == 0 {
					text = text + "posts.Category=\"" + categories[i] + "\""
				} else {
					text = text + " OR posts.Category=\"" + categories[i] + "\""
				}
			}
			text = text + ")"
		} else {
			text = text + " OR "
			for i := range categories {
				// text = text + " OR posts.Category=\"" + categories[i] + "\""
				if i == 0 {
					text = text + "posts.Category=\"" + categories[i] + "\""
				} else {
					text = text + " OR posts.Category=\"" + categories[i] + "\""
				}
			}

		}
	}

	db, err = sql.Open("sqlite3", "./sql/database.db")
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}

	rows, err := db.Query(text)
	defer db.Close()
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	var title string
	var t string
	var n string
	var c string
	var i int
	var likes int
	var dislikes int
	var posts []internal.Post
	var ids []int

	for rows.Next() {
		rows.Scan(&title, &t, &n, &c, &i)
		x := false
		for _, el := range ids {
			if el == i {
				x = true
				break
			}
		}
		if x == true {
			continue
		}
		ids = append(ids, i)

		err := db.QueryRow("SELECT count(*) FROM likes WHERE Postid=(?)", i).Scan(&likes)
		if err != nil {
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		err = db.QueryRow("SELECT count(*) FROM dislikes WHERE Postid=(?)", i).Scan(&dislikes)
		if err != nil {
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		onepost := internal.Post{
			Title:    title,
			Text:     t,
			Name:     n,
			Category: c,
			Id:       i,
			Likes:    likes,
			Dislikes: dislikes,
		}
		posts = append(posts, onepost)

	}

	db.Close()

	if len(r.Form["YourPosts"]) == 1 && (len(categories) != 0 || len(likesdislikes) != 0) {
		fmt.Println("G")
		db, err := sql.Open("sqlite3", "./sql/database.db")
		if err != nil {
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		st, err := db.Query("SELECT lame FROM cookies WHERE Id=(?)", cookie.Value)
		var name string

		for st.Next() {
			st.Scan(&name)
		}
		st.Close()
		var res []internal.Post
		for i := range posts {
			if posts[i].Name == name {
				res = append(res, posts[i])
			}
		}
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
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, res)
		return

	} else if len(r.Form["YourPosts"]) == 1 && len(categories) == 0 && len(likesdislikes) == 0 {
		fmt.Println("GG")
		db, err := sql.Open("sqlite3", "./sql/database.db")
		if err != nil {
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		st, err := db.Query("SELECT lame FROM cookies WHERE Id=(?)", cookie.Value)
		var name string

		for st.Next() {
			st.Scan(&name)
		}

		st.Close()
		db, err = sql.Open("sqlite3", "./sql/database.db")
		if err != nil {
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		var res1 []internal.Post
		st1, err := db.Query("SELECT Title, Post,Namae,Category,Id FROM posts WHERE Namae=(?)", name)
		if err != nil {
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		var title string
		var t string
		var n string
		var c string
		var i int
		var likes int
		var dislikes int
		defer st1.Close()
		for st1.Next() {
			st1.Scan(&title, &t, &n, &c, &i)

			err := db.QueryRow("SELECT count(*) FROM likes WHERE Postid=(?)", i).Scan(&likes)
			if err != nil {
				ErrorHandler(w, http.StatusInternalServerError)
				return
			}
			err = db.QueryRow("SELECT count(*) FROM dislikes WHERE Postid=(?)", i).Scan(&dislikes)
			if err != nil {
				ErrorHandler(w, http.StatusInternalServerError)
				return
			}
			onepost := internal.Post{
				Title:    title,
				Text:     t,
				Name:     n,
				Category: c,
				Id:       i,
				Likes:    likes,
				Dislikes: dislikes,
			}
			res1 = append(res1, onepost)

		}
		db.Close()

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
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, res1)

	} else {

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
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, posts)
	}
}
