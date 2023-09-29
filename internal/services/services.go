package services

type Service struct {
	UserService    UserService
	PostService    PostService
	CommentService CommentService
}

type (
	UserService    interface{}
	PostService    interface{}
	CommentService interface{}
)
