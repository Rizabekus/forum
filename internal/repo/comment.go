package repo

import (
	"database/sql"
	"fmt"
	"forum/internal/models"
	"log"
)

func AddComment(name, text string, id int, db *sql.DB) {
	i, err := db.models
	var count int
	defer i.Close()
	for i.Next() {
		i.Scan(&count)
	}
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	count1, err := db.Query("SELECT count(*) FROM comments WHERE Id=(?)", id)
	if err != nil {
		log.Fatal(err)
	}
	var comid int
	defer count1.Close()
	for count1.Next() {
		count1.Scan(&comid)
	}

	_, err = db.Exec("INSERT INTO comments (Name,Text,Id, Comid) VALUES (?, ?, ?, ?)", name, text, id, comid+1)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
	db.Close()
}

func CollectComments(id int, db *sql.DB) []models.Comment {
	var result []models.Comment
	var name string
	var text string
	st, err := db.Query("SELECT Name, Text, Comid FROM comments WHERE Id=(?)", id)
	if err != nil {
		log.Fatal(err)
	}
	var likes int
	var dislikes int
	var comid int
	for st.Next() {
		st.Scan(&name, &text, &comid)
		err := db.QueryRow("SELECT count(*) FROM comlikes WHERE (Comid,Id)=( ?, ? )", comid, id).Scan(&likes)
		if err != nil {
			log.Fatal(err)
		}
		err = db.QueryRow("SELECT count(*) FROM comdislikes WHERE (Comid,Id)=(? , ? )", comid, id).Scan(&dislikes)
		if err != nil {
			log.Fatal(err)
		}
		x := Comment{
			Name:     name,
			Text:     text,
			Comid:    comid,
			Likes:    likes,
			Dislikes: dislikes,
		}
		result = append(result, x)
	}
	fmt.Println(result)
	return result
}
