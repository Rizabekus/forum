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

func (CookiesService *CookiesService) GetCookie(r *http.Request) *http.Cookie {
	cookie, err := r.Cookie("logged-in")
	if err != nil {
		log.Fatal(err)
	}
	return cookie
}

func (CookiesService *CookiesService) SetCookie(w http.ResponseWriter, cookie *http.Cookie) {
	http.SetCookie(w, cookie)
}

func (CookiesService *CookiesService) DeleteCookie(cookie string) {
	CookiesService.repo.DeleteCookie(cookie)
}
