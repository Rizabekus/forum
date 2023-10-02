package repo

import (
	"database/sql"
	"forum/internal/models"
	"log"
)

type PostDB struct {
	DB *sql.DB
}

func CreatePostRepository(db *sql.DB) *PostDB {
	return &PostDB{DB: db}
}

func (db *PostDB) ShowPost() []models.Post {
	var posts []models.Post

	tx, err := db.DB.Begin()
	if err != nil {
		log.Fatal(err)
	}
	row, err := db.DB.Query("SELECT * FROM posts")
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
		err := db.DB.QueryRow("SELECT count(*) FROM likes WHERE Postid=(?)", i).Scan(&likes)
		if err != nil {
			log.Fatal(err)
		}
		err = db.DB.QueryRow("SELECT count(*) FROM dislikes WHERE Postid=(?)", i).Scan(&dislikes)
		if err != nil {
			log.Fatal(err)
		}
		onepost := models.Post{
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
	db.DB.Close()
	return posts
}

func (db *PostDB) CreatePost(cookie string, text string, category string, title string) {
	Name, err := db.DB.Query("SELECT lame FROM cookies WHERE Id = ( ? )", cookie)
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
	tx, err := db.DB.Begin()
	if err != nil {
		log.Fatal(err)
	}
	Flag, err := db.DB.Query("SELECT count(*) FROM posts")
	defer Flag.Close()
	var flag int
	for Flag.Next() {
		Flag.Scan(&flag)
	}

	_, err = db.DB.Exec("INSERT INTO posts (Title,Post,Namae,Category,Id) VALUES (?, ?, ?, ?, ? )", title, text, name, category, flag+1)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
	db.DB.Close()
}
