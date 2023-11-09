package models

import "net/http"

type Postpage struct {
	Title    string
	Post     string
	Name     string
	Comments []Comment
	Image    string
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
	Image    string
}

type User struct {
	User           string
	Email          string
	HashedPassword string
}

type (
	UserService interface {
		AddUser(UserName string, Email string, hashedPassword string)
		ConfirmSignup(Name string, Email string, Password string, RewrittenPassword string) (bool, string)
		ConfirmSignin(Name string, Password string) (bool, string)
		CreateSession(id string, name string)
		FindUserByToken(cookie string) string
		CheckUserLogin(r *http.Request) bool
	}
	PostService interface {
		ShowPost() []Post
		CreatePost(cookie string, text string, category string, title string, image []byte)
		CountPosts() int
		SelectPostByID(id int) (string, string, string, string)
		LikePost(user string, id string)
		DislikePost(user string, id string)
		Filter(namecookie string, likesdislikes []string, categories []string, yourposts []string) []Post
	}
	CommentService interface {
		AddComment(name, text string, id int)
		CollectComments(id int) []Comment
		LikeComment(checkname string, id string, postid string)
		DislikeComment(checkname string, id string, postid string)
	}

	CookiesService interface {
		SetCookie(w http.ResponseWriter, cookie *http.Cookie)
		GetCookie(r *http.Request) *http.Cookie
		DeleteCookie(cookie string)
	}
)

type (
	UserRepository interface {
		AddUser(UserName string, Email string, hashedPassword string)
		ConfirmSignup(Name string, Email string, Password string, RewrittenPassword string) (bool, string)
		ConfirmSignin(Name string, Password string) (bool, string)
		CreateSession(id string, name string)
		FindUserByToken(cookie string) string
	}
	PostRepository interface {
		ShowPost() []Post
		CreatePost(cookie string, text string, category string, title string, image []byte)
		CountPosts() int
		SelectPostByID(id int) (string, string, string, string)
		RemoveDislikeAtPost(user string, id string)
		RemoveLikeAtPost(user string, id string)
		AddLikeToPost(user string, id string)
		AddDislikeToPost(user string, id string)
		PostDislikeExistence(user string, id string) bool
		PostLikeExistence(user string, id string) bool
		Filter(namecookie string, likesdislikes []string, categories []string, yourposts []string, text string) []Post
	}
	CommentRepository interface {
		AddComment(name, text string, id int)
		CollectComments(id int) []Comment
		CommentDislikeExistence(checkname string, id string, postid string) bool
		CommentLikeExistence(checkname string, id string, postid string) bool
		AddLikeToComment(checkname string, id string, postid string)
		RemoveDislikeFromComment(checkname string, id string, postid string)
		AddDislikeikeToComment(checkname string, id string, postid string)
		RemoveLikeFromComment(checkname string, id string, postid string)
	}

	CookiesRepository interface {
		DeleteCookie(cookie string)
	}
)
