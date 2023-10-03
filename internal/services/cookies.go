package services

import (
	"forum/internal/models"
	"log"
	"net/http"
)

type CookiesService struct {
	repo models.CookiesRepository
}

func CreateCookiesService(repo models.CookiesRepository) *CookiesService {
	return &CookiesService{repo: repo}
}

func (CookieService *CookiesService) GetCookie(r *http.Request) *http.Cookie {
	cookie, err := r.Cookie("logged-in")
	if err != nil {
		log.Fatal(err)
	}
	return cookie
}

func (CookieService *CookiesService) SetCookie(w http.ResponseWriter, cookie *http.Cookie) {
	http.SetCookie(w, cookie)
}

func (CookieService *CookiesService) DeleteCookie(cookie string) {
	CookieService.repo.DeleteCookie(cookie)
}
