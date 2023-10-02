package services

import "forum/internal/models"

type PostService struct {
	repo models.PostRepository
}

func CreatePostService(repo models.PostRepository) *PostService {
	return &PostService{repo: repo}
}

func (PostService *PostService) ShowPost() []models.Post {
	return PostService.repo.ShowPost()
}
