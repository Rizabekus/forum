package services

import (
	"fmt"
	"forum/internal/models"
	"net/http"
)

type CommentService struct {
	repo models.CommentRepository
}

func CreateCommentService(repo models.CommentRepository) *CommentService {
	return &CommentService{repo: repo}
}

func (CommentService *CommentService) AddComment(name, text string, id int) {
	CommentService.repo.AddComment(name, text, id)
}

func (CommentService *CommentService) CollectComments(id int) []models.Comment {
	return CommentService.repo.CollectComments(id)
}

func (CommentService *CommentService) LikeComment() {
	checklikes := false
	checkdislikes := false
	rows, err := db.Query("SELECT Name FROM comlikes WHERE (Comid,Id)=(?,?)", id, postid)
	if err != nil {
		controllers.ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	var likerName string
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&likerName)
		if likerName == checkName {
			checklikes = true
		}
	}
	x, err := db.Query("SELECT Name FROM comdislikes WHERE (Comid,Id)=(?,?)", id, postid)
	if err != nil {
		controllers.ErrorHandler(w, http.StatusInternalServerError)
		return
	}

	var dislikerName string
	defer x.Close()

	for x.Next() {
		x.Scan(&dislikerName)
		if dislikerName == checkName {
			checkdislikes = true
		}
	}
	tx, err := db.Begin()
	if err != nil {
		controllers.ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	fmt.Println("GG", id)
	if checklikes == false && checkdislikes == true {

		_, err = db.Exec("INSERT INTO comlikes (Name, Comid,Id) VALUES (?, ?, ?)", checkName, id, postid)
		_, err = db.Exec("DELETE FROM comdislikes WHERE Name=(?) and Comid=(?) and Id=(?)", checkName, id, postid)
		if err != nil {
			controllers.ErrorHandler(w, http.StatusInternalServerError)
			return
		}

	} else if checklikes == false && checkdislikes == false {

		_, err = db.Exec("INSERT INTO comlikes (Name, Comid,Id) VALUES (?, ?, ?)", checkName, id, postid)

		if err != nil {
			controllers.ErrorHandler(w, http.StatusInternalServerError)
			return
		}

	} else if checklikes == true && checkdislikes == false {

		_, err = db.Exec("DELETE FROM comlikes WHERE Name=(?) and Comid=(?) and Id=(?)", checkName, id, postid)

		if err != nil {
			controllers.ErrorHandler(w, http.StatusInternalServerError)
			return
		}
	}
	tx.Commit()
}
