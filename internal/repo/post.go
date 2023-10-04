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

func (db *PostDB) CountPosts() int {
	count, err := db.DB.Query("select count(*) from posts;")
	if err != nil {
		log.Fatal(err)
	}
	var i int
	defer count.Close()
	for count.Next() {
		count.Scan(&i)
	}
	return i
}

func (db *PostDB) SelectPostByID(id int) (string, string, string) {
	qu, err := db.DB.Query("select Title, Post,Namae from posts where Id=(?)", id)
	if err != nil {
		log.Fatal(err)
	}
	defer qu.Close()
	var title string
	var text string
	var name string
	for qu.Next() {
		qu.Scan(&title, &text, &name)
	}
	return title, text, name
	// db.Close()
}

func PostLikeExistence(db *sql.DB) *likesdislikesDB {
	return &likesdislikesDB{DB: db}
}

func (db *PostDB) PostLikeExistence(user string, id string) bool {
	checklikes := false
	rows, err := db.DB.Query("SELECT Name FROM likes WHERE Postid=(?)", id)
	if err != nil {
		log.Fatal(err)
	}
	var likerName string
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&likerName)
		if likerName == user {
			checklikes = true
		}
	}
	return checklikes
}

func (db *PostDB) PostDislikeExistence(user string, id string) bool {
	checkdislikes := false
	x, err := db.DB.Query("SELECT Name FROM dislikes WHERE Postid=(?)", id)
	if err != nil {
		log.Fatal(err)
	}
	var dislikerName string
	defer x.Close()

	for x.Next() {
		x.Scan(&dislikerName)
		if dislikerName == user {
			checkdislikes = true
		}
	}
	return checkdislikes
}

func (db *PostDB) AddLikeToPost(user string, id string) {
	tx, err := db.DB.Begin()
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.DB.Exec("INSERT INTO likes (Name, Postid) VALUES (?, ?)", user, id)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
}

func (db *PostDB) AddDislikeToPost(user string, id string) {
	tx, err := db.DB.Begin()
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.DB.Exec("INSERT INTO dislikes (Name, Postid) VALUES (?, ?)", user, id)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
}

func (db *PostDB) RemoveLikeAtPost(user string, id string) {
	tx, err := db.DB.Begin()
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.DB.Exec("DELETE FROM likes WHERE Name=(?) and Postid=(?)", user, id)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
}

func (db *PostDB) RemoveDislikeAtPost(user string, id string) {
	tx, err := db.DB.Begin()
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.DB.Exec("DELETE FROM dislikes WHERE Name=(?) and Postid=(?)", user, id)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
}
