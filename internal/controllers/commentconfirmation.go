package controllers

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"
)

func (controllers *Controllers) CommentConfirmation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
		return
	}
	cookie, err := r.Cookie("logged-in")
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	if cookie.Value == "not-logged" {
		http.Redirect(w, r, "http://127.0.0.1:8000/signin?", 302)
	} else {
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
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}

		st, err := db.Query("SELECT lame FROM cookies WHERE Id=(?)", cookie.Value)
		if err != nil {
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		var name string
		for st.Next() {
			st.Scan(&name)
		}
		st.Close()

		internal.AddComment(name, text, id, db)
		http.Redirect(w, r, previousURL, 302)
	}
}
