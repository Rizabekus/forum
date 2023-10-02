package pkg

import (
	"database/sql"
	"log"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func PostChecker(title string, text string) (bool, string) {
	if len(title) < 2 {
		return false, "Too short for title."
	} else if len(text) < 2 {
		return false, "Too short for post."
	}
	if len(title) > 40 {
		return false, "Too long for title."
	}
	// rtitle := []rune(title)
	// rtext := []rune(text)
	// for _, el := range rtitle {
	// 	if el >= 126 || el <= 32 {
	// 		return false, "Title needs to have ASCII symbols only."
	// 	}
	// }
	// for _, el := range rtext {
	// 	if el >= 126 || el <= 32 {
	// 		return false, "Post text needs to have ASCII symbols only."
	// 	}
	// }
	return true, "OK"
}

func ConfirmSignin(Name string, Password string) (bool, string) {
	db, _ := sql.Open("sqlite3", "./sql/database.db")

	rows, err := db.Query("SELECT Name,Password FROM users")
	if err != nil {
		log.Fatal(err)
	}
	var name string
	var password string
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&name, &password)
		if name == Name {
			if bcrypt.CompareHashAndPassword([]byte(password), []byte(Password)) == nil {
				return true, "OK"
			} else {
				return false, "Wrong Password."
			}
		}
	}

	return false, "User does not exist."
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRegex.MatchString(e)
}
