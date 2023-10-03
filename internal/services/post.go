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

func (PostService *PostService) CountPosts() int {
	return PostService.repo.CountPosts()
}

func (PostService *PostService) SelectPostByID(id int) (string, string, string) {
	return PostService.repo.SelectPostByID(id)
}
