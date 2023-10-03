package repo

import (
	"database/sql"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type UserDB struct {
	DB *sql.DB
}

func CreateUserRepository(db *sql.DB) *UserDB {
	return &UserDB{DB: db}
}

func (UserDB *UserDB) AddUser(UserName string, Email string, hashedPassword string) {
	statement, err := UserDB.DB.Prepare("INSERT INTO users (Name, Email,Password) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	statement.Exec(UserName, Email, hashedPassword)
	UserDB.DB.Close()
}

func (UserDB *UserDB) CreateSession(id, name string) {
	tx, err := UserDB.DB.Begin()
	if err != nil {
		log.Fatal(err)
	}

	_, err = UserDB.DB.Exec("INSERT INTO cookies (Id, lame) VALUES (?, ?)", id, name)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
	UserDB.DB.Close()
}

func (UserDB *UserDB) ConfirmSignup(Name string, Email string, Password string, RewrittenPassword string) (bool, string) {
	query := fmt.Sprintf("SELECT Name FROM users")
	row1 := UserDB.DB.QueryRow(query)

	var name string
	err := row1.Scan(&name)

	if err == sql.ErrNoRows {
	} else if err != nil {
		log.Fatal(err)
	} else {

		rows, err := UserDB.DB.Query("SELECT Name, Email FROM users")
		if err != nil {
			log.Fatal(err)
		}
		var name string
		var email string
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&name, &email)
			if name == Name {
				return false, "That name is already being used."
			} else if Email == email {
				return false, "That Email is already being used."
			}
		}

	}

	return true, "OK"
}

func (UserDB *UserDB) ConfirmSignin(Name string, Password string) (bool, string) {
	rows, err := UserDB.DB.Query("SELECT Name,Password FROM users")
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

func (UserDB *UserDB) FindUserByToken(cookie string) string {
	st, err := UserDB.DB.Query("SELECT lame FROM cookies WHERE Id=(?)", cookie)
	if err != nil {
		log.Fatal(err)
	}
	var name string
	for st.Next() {
		st.Scan(&name)
	}
	st.Close()
	return name
}
