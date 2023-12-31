package services

import (
	"forum/internal/models"
	"forum/internal/repo"
)

type Services struct {
	UserService    models.UserService
	PostService    models.PostService
	CommentService models.CommentService
	CookiesService models.CookiesService
}

func ServiceInstance(repo *repo.Repository) *Services {
	return &Services{
		UserService:    CreateUserService(repo.UserRepository),
		PostService:    CreatePostService(repo.PostRepository),
		CommentService: CreateCommentService(repo.CommentRepository),
		CookiesService: CreateCookiesService(repo.CookiesRepository),
	}
}
