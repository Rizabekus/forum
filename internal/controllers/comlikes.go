package controllers

import (
	"net/http"
	"strings"
)

func (controllers *Controllers) ComLikes(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		controllers.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	previousURL := r.Header.Get("Referer")
	postid := (strings.Split(previousURL, "id="))[1]
	id := r.FormValue("id")
	username := controllers.Service.UserService.FindUserByToken(controllers.Service.CookiesService.GetCookie(r).Value)
	controllers.Service.CommentService.LikeComment(username, id, postid)

	http.Redirect(w, r, previousURL, 302)
}
