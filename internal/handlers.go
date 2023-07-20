package internal

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/gofrs/uuid"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

func Homepage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
		return
	}
	if r.URL.Path != "/" {
		http.NotFound(w, r)

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
			log.Println(err.Error())
			return
		}
		tmpl.Execute(w, ShowPost())
	} else if err != nil {
		log.Fatal(err)
	} else {
		c := cookie.Value
		db, err := sql.Open("sqlite3", "./sql/database.db")
		var name string

		Name, err := db.Query("SELECT lame FROM cookies WHERE Id = ( ? )", c)
		if err != nil {
			log.Fatal(err)
		}
		defer Name.Close()
		for Name.Next() {
			Name.Scan(&name)
			fmt.Println(name)
		}

		files := []string{
			"./ui/html/user.home.tmpl",
			"./ui/html/base.layout.tmpl",
		}
		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			log.Println(err.Error())
			return
		}
		db.Close()
		tmpl.Execute(w, ShowPost())

	}
}

// 	if err != nil {

// 		log.Fatal(err)
// 		return
// 	}
// 	var name string
// 	Name.Scan(&name)

// 	if r.URL.Path != "/" {
// 		http.NotFound(w, r)

// 		return
// 	}
// 	if name == "" {
// 		if r.Method != http.MethodGet {
// 			w.Header().Set("Allow", http.MethodGet)
// 			w.WriteHeader(http.StatusMethodNotAllowed)
// 			w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
// 			return
// 		}
// 		files := []string{
// 			"./ui/html/home.page.tmpl",
// 			"./ui/html/base.layout.tmpl",
// 		}
// 		tmpl, err := template.ParseFiles(files...)
// 		if err != nil {
// 			log.Println(err.Error())
// 			return
// 		}
// 		tmpl.Execute(w, nil)
// 	} else {
// 		files := []string{
// 			"./ui/html/user.home.tmpl",
// 			"./ui/html/base.layout.tmpl",
// 		}
// 		tmpl, err := template.ParseFiles(files...)
// 		if err != nil {
// 			log.Println(err.Error())
// 			return
// 		}
// 		fmt.Println(cookie.Value)

// 		tmpl.Execute(w, name)
// 	}
// }

func SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
		return
	}
	tmpl, err := template.ParseFiles("./ui/html/signup.html")
	if err != nil {
		log.Println(err.Error())
		return
	}
	tmpl.Execute(w, nil)
}

func SignUpConfirmation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
		return
	}
	name := r.FormValue("UserName")

	email := r.FormValue("UserEmail")
	password := r.FormValue("UserPassword")
	rewrittenPassword := r.FormValue("UserRewrittenPassword")

	result, text := ConfirmSignup(name, email, password, rewrittenPassword)
	if result == true {

		pwd, err := bcrypt.GenerateFromPassword([]byte(password), 1)
		if err != nil {
			log.Fatal(err)
		}
		db, err := sql.Open("sqlite3", "./sql/database.db")
		if err != nil {
			log.Fatal(err)
		}
		AddUser(name, email, string(pwd), db)

		http.Redirect(w, r, "/signin", 302)
	} else {
		tmpl, err := template.ParseFiles("./ui/html/signup.html")
		if err != nil {
			log.Println(err.Error())
			return
		}
		tmpl.Execute(w, text)
	}
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
		return
	}
	tmpl, err := template.ParseFiles("./ui/html/signin.html")
	if err != nil {
		log.Println(err.Error())
		return
	}
	tmpl.Execute(w, nil)
}

func SignInConfirmation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
		return
	}
	name := r.FormValue("UserName")
	password := r.FormValue("UserPassword")
	result, text := ConfirmSignin(name, password)
	if result == true {
		u1, err := uuid.NewV4()
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		u2 := uuid.NewV3(u1, name).String()
		db, err := sql.Open("sqlite3", "./sql/database.db")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		CreateSession(u2, name, db)

		cookie := &http.Cookie{Name: "logged-in", Value: u2, Expires: time.Now().Add(365 * 24 * time.Hour)}
		http.SetCookie(w, cookie)
		http.Redirect(w, r, "/", 302)
	} else {
		tmpl, err := template.ParseFiles("./ui/html/signin.html")
		if err != nil {
			log.Println(err.Error())
			return
		}
		tmpl.Execute(w, text)
	}
}

func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
		return
	}

	tmpl, err := template.ParseFiles("./ui/html/create.html")
	if err != nil {
		log.Println(err.Error())
		return
	}
	tmpl.Execute(w, nil)
}

func PostConfirmation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
		return
	}
	text := r.FormValue("convert")
	cat := r.FormValue("cars")
	cookie, err := r.Cookie("logged-in")
	if err != nil {
		log.Fatal(err)
	}
	CreatePost(cookie.Value, text, cat)
	http.Redirect(w, r, "/", 302)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./sql/database.db")
	if err != nil {
		log.Fatal(err)
	}
	cookie, err := r.Cookie("logged-in")
	DeleteCookie(cookie.Value, db)
	cookie = &http.Cookie{
		Name:  "logged-in",
		Value: "not-logged",
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", 302)
}

func PostPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
		return
	}
	xurl := strings.Split(r.URL.String(), "id=")
	if len(xurl) < 2 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	id, err := strconv.Atoi(xurl[1])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	if r.URL.String() != "/comments?id="+strconv.Itoa(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	db, err := sql.Open("sqlite3", "./sql/database.db")
	defer db.Close()
	count, err := db.Query("select count(*) from posts;")
	if err != nil {
		log.Fatal(err)
	}
	var i int
	defer count.Close()
	for count.Next() {
		count.Scan(&i)
	}
	if id > i || id < 1 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	tmp, err := template.ParseFiles("./ui/html/comments.html")
	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = tmp.Execute(w, nil)
}

func CommentConfirmation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
		return
	}
	text := r.FormValue("comment")
	previousURL := r.Header.Get("Referer")

	xurl := strings.Split(previousURL, "id=")
	if len(xurl) < 2 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	id, err := strconv.Atoi(xurl[1])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	db, err := sql.Open("sqlite3", "./sql/database.db")
	if err != nil {
		log.Fatal(err)
	}
	st, err := db.Query("SELECT Namae FROM posts WHERE Id=(?)", id)
	if err != nil {
		log.Fatal(err)
	}
	var name string
	for st.Next() {
		st.Scan(&name)
	}
	AddComment(name, text, id, db)
	http.Redirect(w, r, previousURL, 302)
}
