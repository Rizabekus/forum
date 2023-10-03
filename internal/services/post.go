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

func (PostService *PostService) CreatePost(cookie string, text string, category string, title string) {
	PostService.repo.CreatePost(cookie, text, category, title)
}
