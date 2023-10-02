package services

import "forum/internal/models"

type UserService struct {
	repo models.UserRepository
}

func CreateUserService(repo models.UserRepository) *UserService {
	return &UserService{repo: repo}
}
