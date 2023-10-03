package services

import "forum/internal/models"

type CommentService struct {
	repo models.CommentRepository
}

func CreateCommentService(repo models.CommentRepository) *CommentService {
	return &CommentService{repo: repo}
}

func (CommentService *CommentService) AddComment(name, text string, id int) {
	CommentService.repo.AddComment(name, text, id)
}
