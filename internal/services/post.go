package services

import "forum/internal/models"

type PostService struct {
	repo models.CommentRepository
}

func CreatePostService(repo models.PostRepository) *PostService {
	return &PostService{repo: repo}
}
