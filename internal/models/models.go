package models

type Postpage struct {
	Title    string
	Post     string
	Name     string
	Comments []Comment
}

type ErrorStruct struct {
	Status  int
	Message string
}

type UserService interface{}

type (
	PostService          interface{}
	CommentService       interface{}
	LikesDislikesService interface{}
)
