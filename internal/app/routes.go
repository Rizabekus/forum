package app

import (
	"fmt"
	"forum/internal/controllers"
	"log"
	"net/http"
)

func Routes(c *controllers.Controllers) {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("./templates"))
	mux.Handle("/templates/", http.StripPrefix("/templates", files))

	mux.HandleFunc("/", c.Homepage)
	mux.HandleFunc("/signup", c.SignUp)
	mux.HandleFunc("/signin", c.SignIn)
	mux.HandleFunc("/logout", c.Logout)
	mux.HandleFunc("/signupconfirmation", c.SignUpConfirmation)
	mux.HandleFunc("/signinconfirmation", c.SignInConfirmation)
	mux.HandleFunc("/comments", c.PostPage)
	mux.HandleFunc("/postconfirmation", c.PostConfirmation)
	mux.HandleFunc("/commentconfirmation", c.CommentConfirmation)
	mux.HandleFunc("/create", c.CreatePost)
	mux.HandleFunc("/like", c.Likes)
	mux.HandleFunc("/dislike", c.Dislikes)
	mux.HandleFunc("/filter", c.Filter)
	mux.HandleFunc("/comlike", c.ComLikes)
	mux.HandleFunc("/comdislike", c.ComDislikes)
	mux.HandleFunc("/filter/likes", c.Likes)
	fmt.Println("http://localhost:8000")

	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)
}
