package auth

import (
	"../models"
	"../utils/database"
	"../utils/password"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func userExist(username string) bool {
	db := database.Connect()
	defer db.Close()
	stmt, err := db.Prepare("SELECT * FROM users WHERE username=?")

	if err != nil {
		fmt.Println(err)
		return false
	}
	defer stmt.Close()

	rows, err := stmt.Query(username)
	if err != nil {
		fmt.Println(err)
		return false
	}

	defer rows.Close()

	if !rows.Next() {
		fmt.Println("user not found, user with this username doesn't exit, you can create it")
		return true
	}
	return false
}

func emailExist(email string) bool {
	db := database.Connect()
	defer db.Close()
	stmt, err := db.Prepare("SELECT * FROM users WHERE email=?")

	if err != nil {
		fmt.Println(err)
		return false
	}
	defer stmt.Close()

	rows, err := stmt.Query(email)
	if err != nil {
		fmt.Println(err)
		return false
	}

	defer rows.Close()

	if !rows.Next() {
		fmt.Println("user with this email doesn't exit, you can create it")
		return true
	}
	return false
}

func Register(username string, email string, pass string) (bool, string) {
	newUserName, newUserEmail := userExist(username), emailExist(email)

	if newUserName && newUserEmail {
		db := database.Connect()
		defer db.Close()
		hash, err := password.HashPassword(pass)
		if err != nil {
			return false, "hash err"
		}
		database.AddUser(db, models.User{UserName: username, Email: email, Password: hash})

		user := models.User{UserName: username, Email: email, Password: pass}
		tokenContent := jwt.MapClaims{
			"user_id": user.ID,
			"expiry":  time.Now().Add(time.Minute * 60).Unix(),
		}
		jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)

		token, err := jwtToken.SignedString([]byte("TokenPassword"))

		if err != nil {
			panic(err.Error())
		}
		return true, token
	}
	return false, ""
}