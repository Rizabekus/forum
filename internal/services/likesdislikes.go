package services

import "forum/internal/models"

type LikesDislikesService struct {
	repo models.LikesDislikesRepository
}

func CreateLikesDislikesService(repo models.LikesDislikesRepository) *LikesDislikesService {
	return &LikesDislikesService{repo: repo}
}
