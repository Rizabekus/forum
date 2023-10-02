package repo

import (
	"database/sql"
	"forum/internal/models"

	_ "github.com/mattn/go-sqlite3"
)

func RepositoryInstance(db *sql.DB) *Repository {
	return &Repository{
		UserRepository:          CreateUserRepository(db),
		PostRepository:          CreatePostRepository(db),
		CommentsReepository:     CreateCommentsRepository(db),
		LikesdislikesRepository: CreateLikesdislikesRepository(db),
	}
}

type Repository struct {
	UserRepository          models.UserRepository
	PostRepository          models.PostRepository
	CommentsReepository     models.CommentsReepository
	LikesdislikesRepository models.LikesdislikesRepository
}

// ype Post struct {
// 	Title    string
// 	Text     string
// 	Name     string
// 	Category string
// 	Id       int
// 	Likes    int
// 	Dislikes int
