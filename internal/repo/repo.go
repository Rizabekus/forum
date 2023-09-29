package repo

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func RepositoryInstance(db *sql.DB) *Repository {
	return &Repository{
		UserRepository:          UserRepository,
		PostRepository:          PostRepository,
		CommentsReepository:     CommentsReepository,
		LikesdislikesRepository: LikesdislikesRepository,
	}
}

type Repository struct {
	UserRepository          UserRepository
	PostRepository          PostRepository
	CommentsReepository     CommentsReepository
	LikesdislikesRepository LikesdislikesRepository
}

type (
	UserRepository          interface{}
	PostRepository          interface{}
	CommentsReepository     interface{}
	LikesdislikesRepository interface{}
)

// ype Post struct {
// 	Title    string
// 	Text     string
// 	Name     string
// 	Category string
// 	Id       int
// 	Likes    int
// 	Dislikes int
