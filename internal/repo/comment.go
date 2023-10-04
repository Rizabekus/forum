package repo

import (
	"database/sql"
	"forum/internal/models"
	"log"
)

type CommentDB struct {
	DB *sql.DB
}

func CreateCommentRepository(db *sql.DB) *CommentDB {
	return &CommentDB{DB: db}
}

func (db *CommentDB) AddComment(name, text string, id int) {
	i, err := db.DB.Query("SELECT count(*) from comments where id = (?)", id)
	var count int
	defer i.Close()
	for i.Next() {
		i.Scan(&count)
	}
	tx, err := db.DB.Begin()
	if err != nil {
		log.Fatal(err)
	}
	count1, err := db.DB.Query("SELECT count(*) FROM comments WHERE Id=(?)", id)
	if err != nil {
		log.Fatal(err)
	}
	var comid int
	defer count1.Close()
	for count1.Next() {
		count1.Scan(&comid)
	}

	_, err = db.DB.Exec("INSERT INTO comments (Name,Text,Id, Comid) VALUES (?, ?, ?, ?)", name, text, id, comid+1)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
}

func (db *CommentDB) CollectComments(id int) []models.Comment {
	var result []models.Comment
	var name string
	var text string
	st, err := db.DB.Query("SELECT Name, Text, Comid FROM comments WHERE Id=(?)", id)
	if err != nil {
		log.Fatal(err)
	}
	var likes int
	var dislikes int
	var comid int
	for st.Next() {
		st.Scan(&name, &text, &comid)
		err := db.DB.QueryRow("SELECT count(*) FROM comlikes WHERE (Comid,Id)=( ?, ? )", comid, id).Scan(&likes)
		if err != nil {
			log.Fatal(err)
		}
		err = db.DB.QueryRow("SELECT count(*) FROM comdislikes WHERE (Comid,Id)=(? , ? )", comid, id).Scan(&dislikes)
		if err != nil {
			log.Fatal(err)
		}
		x := models.Comment{
			Name:     name,
			Text:     text,
			Comid:    comid,
			Likes:    likes,
			Dislikes: dislikes,
		}
		result = append(result, x)
	}

	return result
}

func (db *CommentDB) CommentLikeExistence(checkname string, id string, postid string) bool {
	checklikes := false
	rows, err := db.DB.Query("SELECT Name FROM comlikes WHERE (Comid,Id)=(?,?)", id, postid)
	if err != nil {
		log.Fatal(err)
	}
	var likerName string
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&likerName)
		if likerName == checkname {
			checklikes = true
		}
	}
	return checklikes
}

func (db *CommentDB) CommentDislikeExistence(checkname string, id string, postid string) bool {
	checkdislikes := false
	x, err := db.DB.Query("SELECT Name FROM comdislikes WHERE (Comid,Id)=(?,?)", id, postid)
	if err != nil {
		log.Fatal(err)
	}

	var dislikerName string
	defer x.Close()

	for x.Next() {
		x.Scan(&dislikerName)
		if dislikerName == checkname {
			checkdislikes = true
		}
	}
	return checkdislikes
}

func (db *CommentDB) AddLikeToComment(checkname string, id string, postid string) {
	_, err := db.DB.Exec("INSERT INTO comlikes (Name, Comid,Id) VALUES (?, ?, ?)", checkname, id, postid)
	if err != nil {
		log.Fatal(err)
	}
}

func (db *CommentDB) RemoveDislikeFromComment(checkname string, id string, postid string) {
	_, err := db.DB.Exec("DELETE FROM comdislikes WHERE Name=(?) and Comid=(?) and Id=(?)", checkname, id, postid)
	if err != nil {
		log.Fatal(err)
	}
}

func (db *CommentDB) AddDislikeikeToComment(checkname string, id string, postid string) {
	_, err := db.DB.Exec("INSERT INTO comdislikes (Name, Comid,Id) VALUES (?, ?, ?)", checkname, id, postid)
	if err != nil {
		log.Fatal(err)
	}
}

func (db *CommentDB) RemoveLikeFromComment(checkname string, id string, postid string) {
	_, err := db.DB.Exec("DELETE FROM comlikes WHERE Name=(?) and Comid=(?) and Id=(?)", checkname, id, postid)
	if err != nil {
		log.Fatal(err)
	}
}
