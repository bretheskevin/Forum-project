package auth

import (
	"../models"
	"../utils/database"
	"../utils/password"
	"../utils/validation"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func Register(username string, email string, pass string) (bool, string) {
	newUserName, newUserEmail := validation.UserExist(username), validation.EmailExist(email)

	if newUserName && newUserEmail {
		db := database.Connect()
		defer db.Close()
		hash, err := password.HashPassword(pass)
		if err != nil {
			return false, "hash err"
		}
		database.AddUser(models.User{UserName: username, Email: email, Password: hash})

		user, exist := database.GetUserByEmail(email)
		if !exist {
			return false, "user with this email don't exist"
		}
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

	if !newUserName {
		return false, "name"
	}
	if !newUserEmail {
		return false, "email"
	}

	return false, ""
}
