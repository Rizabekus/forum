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
	UserService interface{}
	PostService interface {
		ShowPost() []Post
	}
	CommentService       interface{}
	LikesDislikesService interface{}
)

type (
	UserRepository interface {
		AddUser(UserName string, Email string, hashedPassword string)
	}
	PostRepository interface {
		ShowPost() []Post
		CreatePost(cookie string, text string, category string, title string)
	}
	CommentRepository interface {
		AddComment(name, text string, id int)
		CollectComments(id int) []Comment
	}
	LikesDislikesRepository interface{}
)
