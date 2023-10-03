package controllers

import (
	"net/http"
)

func (controllers *Controllers) Likes(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		controllers.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	id := r.FormValue("id")

	cookie, err := r.Cookie("logged-in")
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	controllers.Service.LikesDislikesService.Like(id, cookie.Value)

	http.Redirect(w, r, previousURL, 302)
}
