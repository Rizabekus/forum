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

func (PostService *PostService) LikePost(user string, id string) {
	checklikes := PostService.repo.PostLikeExistence(user, id)
	checkdislikes := PostService.repo.PostDislikeExistence(user, id)

	if checklikes == false && checkdislikes == true {
		PostService.repo.AddLikeToPost(user, id)
		PostService.repo.RemoveDislikeAtPost(user, id)

	} else if checklikes == false && checkdislikes == false {
		PostService.repo.AddLikeToPost(user, id)
	} else if checklikes == true && checkdislikes == false {
		PostService.repo.RemoveLikeAtPost(user, id)
	}
}

func (PostService *PostService) DislikePost(user string, id string) {
	checklikes := PostService.repo.PostLikeExistence(user, id)
	checkdislikes := PostService.repo.PostDislikeExistence(user, id)

	if checklikes == true && checkdislikes == false {
		PostService.repo.AddDislikeToPost(user, id)
		PostService.repo.RemoveLikeAtPost(user, id)

	} else if checklikes == false && checkdislikes == false {
		PostService.repo.AddDislikeToPost(user, id)
	} else if checklikes == true && checkdislikes == false {
		PostService.repo.RemoveDislikeAtPost(user, id)
	}
}
