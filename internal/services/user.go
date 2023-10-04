package services

import (
	"forum/internal/models"
	"regexp"
)

type UserService struct {
	repo models.UserRepository
}

func CreateUserService(repo models.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (UserService *UserService) AddUser(UserName string, Email string, hashedPassword string) {
	UserService.repo.AddUser(UserName, Email, hashedPassword)
}

func (UserService *UserService) ConfirmSignup(Name string, Email string, Password string, RewrittenPassword string) (bool, string) {
	if RewrittenPassword != Password {
		return false, "Passwords don't match, write again."
	}
	if len(Name) < 3 || len(Password) < 7 {
		return false, "Name/Password doesn't have enough characters. Minimum for name is 3 and for password is 7."
	}
	r := []rune(Name)
	for i := range r {
		if !((r[i] >= 97 && r[i] <= 122) || (r[i] >= 65 && r[i] <= 90)) && i == 0 {
			return false, "First letter should be a alphabetical character."
		}
		if r[i] == ' ' {
			return false, "Can't have space character in the name."
		}
		if r[i] > 122 || r[i] < 33 {
			return false, "Can only have ASCII characters for Username."
		}
	}
	r = []rune(Password)
	for i := range r {
		if r[i] > 122 || r[i] < 33 {
			return false, "Can only have ASCII characters for Password."
		}
	}
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if emailRegex.MatchString(Email) == false {
		return false, "Wrong format for Email."
	}

	return UserService.repo.ConfirmSignup(Name, Email, Password, RewrittenPassword)
}

func (UserService *UserService) ConfirmSignin(Name string, Password string) (bool, string) {
	return UserService.repo.ConfirmSignin(Name, Password)
}

func (UserService *UserService) CreateSession(id, name string) {
	UserService.repo.CreateSession(id, name)
}

func (UserService *UserService) FindUserByToken(cookie string) string {
	return UserService.repo.FindUserByToken(cookie)
}
