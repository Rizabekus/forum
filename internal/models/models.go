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

type Comment struct {
	Name     string
	Text     string
	Comid    int
	Likes    int
	Dislikes int
}

type Post struct {
	Title    string
	Text     string
	Name     string
	Category string
	Id       int
	Likes    int
	Dislikes int
}

type (
	UserService          interface{}
	PostService          interface{}
	CommentService       interface{}
	LikesDislikesService interface{}
)

type (
	UserRepository          interface{}
	PostRepository          interface{}
	CommentRepository       interface{}
	LikesDislikesRepository interface{}
)
