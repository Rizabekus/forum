package services

import (
	"forum/internal/models"
)

type LikesDislikesService struct {
	repo models.LikesDislikesRepository
}

func CreateLikesDislikesService(repo models.LikesDislikesRepository) *LikesDislikesService {
	return &LikesDislikesService{repo: repo}
}

func (LikesDislikesService *LikesDislikesService) Like(user string, id string) {
	checklikes := LikesDislikesService.repo.CheckLikeExistence(user, id)
	checkdislikes := LikesDislikesService.repo.CheckDislikeExistence(user, id)

	if checklikes == false && checkdislikes == true {
		LikesDislikesService.repo.AddLike(user, id)
		LikesDislikesService.repo.RemoveDislike(user, id)

	} else if checklikes == false && checkdislikes == false {
		LikesDislikesService.repo.AddLike(user, id)
	} else if checklikes == true && checkdislikes == false {
		LikesDislikesService.repo.RemoveLike(user, id)
	}
}

func (LikesDislikesService *LikesDislikesService) Dislike(user string, id string) {
	checklikes := LikesDislikesService.repo.CheckLikeExistence(user, id)
	checkdislikes := LikesDislikesService.repo.CheckDislikeExistence(user, id)

	if checklikes == true && checkdislikes == false {
		LikesDislikesService.repo.AddDislike(user, id)
		LikesDislikesService.repo.RemoveLike(user, id)

	} else if checklikes == false && checkdislikes == false {
		LikesDislikesService.repo.AddDislike(user, id)
	} else if checklikes == true && checkdislikes == false {
		LikesDislikesService.repo.RemoveDislike(user, id)
	}
}
