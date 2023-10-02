package repo

import (
	"database/sql"
	"log"
)

type PostDB struct {
	DB *sql.DB
}

func CreatePostRepository(db *sql.DB) *PostDB {
	return &PostDB{DB: db}
}

func ShowPost() []Post {
	var posts []Post
	db, err := sql.Open("sqlite3", "./sql/database.db")
	if err != nil {
		log.Fatal(err)
	}
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	row, err := db.Query("SELECT * FROM posts")
	if err != nil {
		log.Fatal(err)
	}
	var title string
	var t string
	var n string
	var c string
	var i int
	var likes int
	var dislikes int
	defer row.Close()
	for row.Next() {
		row.Scan(&title, &t, &n, &c, &i)
		err := db.QueryRow("SELECT count(*) FROM likes WHERE Postid=(?)", i).Scan(&likes)
		if err != nil {
			log.Fatal(err)
		}
		err = db.QueryRow("SELECT count(*) FROM dislikes WHERE Postid=(?)", i).Scan(&dislikes)
		if err != nil {
			log.Fatal(err)
		}
		onepost := Post{
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

	tx.Commit()
	db.Close()
	return posts
}

func CreatePost(cookie string, text string, category string, title string) {
	db, err := sql.Open("sqlite3", "./sql/database.db")
	Name, err := db.Query("SELECT lame FROM cookies WHERE Id = ( ? )", cookie)
	if err != nil {
		log.Fatal(err)
	}
	defer Name.Close()
	var name string
	for Name.Next() {
		Name.Scan(&name)
	}
	if err != nil {
		log.Fatal(err)
	}
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	Flag, err := db.Query("SELECT count(*) FROM posts")
	defer Flag.Close()
	var flag int
	for Flag.Next() {
		Flag.Scan(&flag)
	}

	_, err = db.Exec("INSERT INTO posts (Title,Post,Namae,Category,Id) VALUES (?, ?, ?, ?, ? )", title, text, name, category, flag+1)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
	db.Close()
}
