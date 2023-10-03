package controllers

import (
	"net/http"
)

func (controllers *Controllers) Logout(w http.ResponseWriter, r *http.Request) {
	cookie := controllers.Service.CookiesService.GetCookie(r)
	controllers.Service.CookiesService.DeleteCookie(cookie.Value)
	cookie = &http.Cookie{
		Name:  "logged-in",
		Value: "not-logged",
	}
	controllers.Service.CookiesService.SetCookie(w, cookie)
	http.Redirect(w, r, "/", 302)
}
