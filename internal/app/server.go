package app

import (
	"fmt"
	"forum/internal/controllers"
	"log"
	"net/http"
)

func Run() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("./templates"))
	mux.Handle("/templates/", http.StripPrefix("/templates", files))

	mux.HandleFunc("/", controllers.Homepage)
	mux.HandleFunc("/signup", controllers.SignUp)
	mux.HandleFunc("/signin", controllers.SignIn)
	mux.HandleFunc("/logout", controllers.Logout)
	mux.HandleFunc("/signupconfirmation", controllers.SignUpConfirmation)
	mux.HandleFunc("/signinconfirmation", controllers.SignInConfirmation)
	mux.HandleFunc("/comments", controllers.PostPage)
	mux.HandleFunc("/postconfirmation", controllers.PostConfirmation)
	mux.HandleFunc("/commentconfirmation", controllers.CommentConfirmation)
	mux.HandleFunc("/create", controllers.Create)
	mux.HandleFunc("/like", controllers.Likes)
	mux.HandleFunc("/dislike", controllers.Dislikes)
	mux.HandleFunc("/filter", controllers.Filter)
	mux.HandleFunc("/comlike", controllers.ComLikes)
	mux.HandleFunc("/comdislike", controllers.ComDislikes)
	mux.HandleFunc("/filter/likes", controllers.Likes)
	fmt.Println("http://127.0.0.1:8000")
	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)
}
