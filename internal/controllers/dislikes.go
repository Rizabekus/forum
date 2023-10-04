package controllers

import (
	"net/http"
)

func (controllers *Controllers) Dislikes(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		controllers.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	id := r.FormValue("id")

	cookie, err := r.Cookie("logged-in")
	if err != nil {
		controllers.ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	controllers.Service.PostService.DislikePost(controllers.Service.UserService.FindUserByToken(cookie.Value), id)
	previousURL := r.Header.Get("Referer")
	http.Redirect(w, r, previousURL, 302)
}
