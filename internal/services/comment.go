package services

import (
	"forum/internal/models"
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

func (CommentService *CommentService) LikeComment(checkname string, id string, postid string) {
	checklikes := CommentService.repo.CommentLikeExistence(checkname, id, postid)
	checkdislikes := CommentService.repo.CommentDislikeExistence(checkname, id, postid)

	if checklikes == false && checkdislikes == true {
		CommentService.repo.AddLikeToComment(checkname, id, postid)
		CommentService.repo.RemoveDislikeFromComment(checkname, id, postid)

	} else if checklikes == false && checkdislikes == false {
		CommentService.repo.AddLikeToComment(checkname, id, postid)
	} else if checklikes == true && checkdislikes == false {
		CommentService.repo.RemoveLikeFromComment(checkname, id, postid)
	}
}

func (CommentService *CommentService) DislikeComment(checkname string, id string, postid string) {
	checklikes := CommentService.repo.CommentLikeExistence(checkname, id, postid)
	checkdislikes := CommentService.repo.CommentDislikeExistence(checkname, id, postid)

	if checklikes == true && checkdislikes == false {
		CommentService.repo.AddDislikeikeToComment(checkname, id, postid)
		CommentService.repo.RemoveLikeFromComment(checkname, id, postid)

	} else if checklikes == false && checkdislikes == false {
		CommentService.repo.AddDislikeikeToComment(checkname, id, postid)
	} else if checklikes == false && checkdislikes == true {
		CommentService.repo.RemoveDislikeFromComment(checkname, id, postid)
	}
}
