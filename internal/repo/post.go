package repo

import (
	"database/sql"
	"encoding/base64"
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

	// tx, err := db.DB.Begin()
	// if err != nil {
	// 	log.Fatal(err)
	// }
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
	var img []byte

	defer row.Close()
	for row.Next() {
		row.Scan(&title, &t, &n, &c, &i, &img)
		err := db.DB.QueryRow("SELECT count(*) FROM likes WHERE Postid=(?)", i).Scan(&likes)
		if err != nil {
			log.Println("Error in ShowPost")
			log.Fatal(err)
		}
		err = db.DB.QueryRow("SELECT count(*) FROM dislikes WHERE Postid=(?)", i).Scan(&dislikes)
		if err != nil {
			log.Println("Error in ShowPost")
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
			Image:    "data:image/png;base64," + base64.StdEncoding.EncodeToString(img),
		}

		posts = append(posts, onepost)
	}

	// tx.Commit()

	return posts
}

func (db *PostDB) CreatePost(cookie string, text string, category string, title string, image []byte) {
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

	_, err = db.DB.Exec("INSERT INTO posts (Title,Post,Namae,Category,Id,Image) VALUES (?, ?, ?, ?, ?, ? )", title, text, name, category, flag+1, image)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
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
	_, err := db.DB.Exec("INSERT INTO likes (Name, Postid) VALUES (?, ?)", user, id)
	if err != nil {
		log.Fatal(err)
	}
}

func (db *PostDB) AddDislikeToPost(user string, id string) {
	_, err := db.DB.Exec("INSERT INTO dislikes (Name, Postid) VALUES (?, ?)", user, id)
	if err != nil {
		log.Fatal(err)
	}
}

func (db *PostDB) RemoveLikeAtPost(user string, id string) {
	_, err := db.DB.Exec("DELETE FROM likes WHERE Name=(?) and Postid=(?)", user, id)
	if err != nil {
		log.Fatal(err)
	}
}

func (db *PostDB) RemoveDislikeAtPost(user string, id string) {
	_, err := db.DB.Exec("DELETE FROM dislikes WHERE Name=(?) and Postid=(?)", user, id)
	if err != nil {
		log.Fatal(err)
	}
}

func (db *PostDB) Filter(namecookie string, likesdislikes []string, categories []string, yourposts []string, text string) []models.Post {
	rows, err := db.DB.Query(text)
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
	var posts []models.Post
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

	if len(yourposts) == 1 && (len(categories) != 0 || len(likesdislikes) != 0) {

		name := namecookie
		var res []models.Post
		for i := range posts {
			if posts[i].Name == name {
				res = append(res, posts[i])
			}
		}

		return res

	} else if len(yourposts) == 1 && len(categories) == 0 && len(likesdislikes) == 0 {

		name := namecookie

		var res1 []models.Post
		st1, err := db.DB.Query("SELECT Title, Post,Namae,Category,Id FROM posts WHERE Namae=(?)", name)
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
		defer st1.Close()
		for st1.Next() {
			st1.Scan(&title, &t, &n, &c, &i)

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
			res1 = append(res1, onepost)

		}

		return res1

	} else {
		return posts
	}
}
