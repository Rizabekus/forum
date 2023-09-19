package models

type Post struct {
	Title    string
	Text     string
	Name     string
	Category string
	Id       int
	Likes    int
	Dislikes int
	// Comments [string]string
}

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
