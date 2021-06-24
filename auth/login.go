package auth

import (
	"../models"
	"../utils/database"
	"../utils/password"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func Login(email string, pass string) (bool, string) {
	db := database.Connect()
	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM users WHERE email=?")

	if err != nil {
		fmt.Println(err)
		return false, ""
	}

	defer stmt.Close()

	rows, err := stmt.Query(email)
	if err != nil {
		fmt.Println(err)
		return false, ""
	}

	defer rows.Close()

	if !rows.Next() {
		fmt.Println("user not found")
		return false, ""
	}

	user := models.User{}
	if err = rows.Scan(&user.ID, &user.UserName, &user.Email, &user.Password, &user.ProfilePictureURL, &user.IsAdmin); err != nil {
		fmt.Println(err)
		return false, ""
	}

	hash := user.Password

	errPass := password.CheckPassword(pass, hash)
	if errPass != nil {
		fmt.Println("Wrong Password")
		return false, ""
	}

	defer db.Close()

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
