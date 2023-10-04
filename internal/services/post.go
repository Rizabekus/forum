package services

import (
	"forum/internal/models"
)

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

func (PostService *PostService) Filter(namecookie string, likesdislikes []string, categories []string, yourposts []string) []models.Post {
	var formattedlikes []string

	for i := range likesdislikes {
		formattedlikes = append(formattedlikes, likesdislikes[i]+"s.Postid")
	}

	text := "SELECT posts.Title, posts.Post, posts.Namae, posts.Category, posts.Id from posts "

	if len(likesdislikes) == 2 {
		text = text + "INNER JOIN likes on posts.Id=likes.Postid INNER JOIN dislikes on posts.Id=dislikes.Postid"
	} else if len(likesdislikes) == 1 {
		if likesdislikes[0] == "like" {
			text = text + "INNER JOIN likes on posts.Id=likes.Postid"
		} else {
			text = text + "INNER JOIN dislikes on posts.Id=dislikes.Postid"
		}
	}
	if len(likesdislikes) > 0 {

		text = text + " WHERE "
		for i := range formattedlikes {
			if i == 0 {
				text = text + "(posts.Id=" + formattedlikes[i] + " AND " + likesdislikes[i] + "s.Name=\"" + namecookie + "\")"
			} else {
				text = text + " OR (posts.Id=" + formattedlikes[i] + " AND " + likesdislikes[i] + "s.Name=\"" + namecookie + "\")"
			}
		}
	} else if len(categories) > 0 {
		text = text + " WHERE "
		for i := range categories {
			if i == 0 {
				text = text + "posts.Category=\"" + categories[i] + "\""
			} else {
				text = text + " OR posts.Category=\"" + categories[i] + "\""
			}
		}
	}
	if len(categories) > 0 {
		if len(likesdislikes) > 0 {
			text = text + " AND ("
			for i := range categories {
				if i == 0 {
					text = text + "posts.Category=\"" + categories[i] + "\""
				} else {
					text = text + " OR posts.Category=\"" + categories[i] + "\""
				}
			}
			text = text + ")"
		} else {
			text = text + " OR "
			for i := range categories {
				if i == 0 {
					text = text + "posts.Category=\"" + categories[i] + "\""
				} else {
					text = text + " OR posts.Category=\"" + categories[i] + "\""
				}
			}

		}
	}
	posts := PostService.repo.Filter(namecookie, likesdislikes, categories, yourposts, text)
	return posts
}
