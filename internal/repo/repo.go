package repo

import (
	"database/sql"
	"forum/internal/models"

	_ "github.com/mattn/go-sqlite3"
)

func RepositoryInstance(db *sql.DB) *Repository {
	return &Repository{
		UserRepository:    CreateUserRepository(db),
		PostRepository:    CreatePostRepository(db),
		CommentRepository: CreateCommentRepository(db),
		CookiesRepository: CreateCookiesRepository(db),
	}
}

type Repository struct {
	UserRepository    models.UserRepository
	PostRepository    models.PostRepository
	CommentRepository models.CommentRepository
	CookiesRepository models.CookiesRepository
}

// ype Post struct {
// 	Title    string
// 	Text     string
// 	Name     string
// 	Category string
// 	Id       int
// 	Likes    int
// 	Dislikes int
