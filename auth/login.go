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

	stmt, err := db.Prepare("SELECT * FROM users WHERE email=?") // prepare stmt

	if err != nil {
		fmt.Println(err)
		return false, ""
	}

	defer stmt.Close()

	rows, err := stmt.Query(email) // check some match, stock it into rows
	if err != nil {
		fmt.Println(err)
		return false, ""
	}

	defer rows.Close()

	if !rows.Next() { // if rows is empty user not found
		fmt.Println("user not found")
		return false, ""
	}

	user := models.User{}
	// scan info
	if err = rows.Scan(&user.ID, &user.UserName, &user.Email, &user.Password, &user.ProfilePictureURL, &user.IsAdmin); err != nil {
		fmt.Println(err)
		return false, ""
	}

	// hash password
	hash := user.Password

	errPass := password.CheckPassword(pass, hash) // check if the password is good
	if errPass != nil {
		fmt.Println("Wrong Password")
		return false, ""
	}

	defer db.Close()

	tokenContent := jwt.MapClaims{ // create token content
		"user_id": user.ID,
		"expiry":  time.Now().Add(time.Minute * 60).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte("TokenPassword")) // create token signed
	if err != nil {
		panic(err.Error())
	}
	return true, token // return token
}
