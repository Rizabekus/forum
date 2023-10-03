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
	UserService interface {
		AddUser(UserName string, Email string, hashedPassword string)
		ConfirmSignup(Name string, Email string, Password string, RewrittenPassword string) (bool, string)
		ConfirmSignin(Name string, Password string) (bool, string)
		CreateSession(id, name string)
		FindUserByToken(cookie string) string
	}
	PostService interface {
		ShowPost() []Post
		CreatePost(cookie string, text string, category string, title string)
		CountPosts() int
		SelectPostByID(id int) (string, string, string)
	}
	CommentService interface {
		AddComment(name, text string, id int)
		CollectComments(id int) []Comment
	}
	LikesDislikesService interface {
		Like(user string, id string)
		Dislike(user string, id string)
	}
)

type (
	UserRepository interface {
		AddUser(UserName string, Email string, hashedPassword string)
		ConfirmSignup(Name string, Email string, Password string, RewrittenPassword string) (bool, string)
		ConfirmSignin(Name string, Password string) (bool, string)
		CreateSession(id, name string)
		FindUserByToken(cookie string) string
	}
	PostRepository interface {
		ShowPost() []Post
		CreatePost(cookie string, text string, category string, title string)
		CountPosts() int
		SelectPostByID(id int) (string, string, string)
	}
	CommentRepository interface {
		AddComment(name, text string, id int)
		CollectComments(id int) []Comment
	}
	LikesDislikesRepository interface {
		CheckLikeExistence(user string, id string) bool
		CheckDislikeExistence(user string, id string) bool
		RemoveLike(user string, id string)
		RemoveDislike(user string, id string)
		AddDislike(user string, id string)
		AddLike(user string, id string)
	}
)
