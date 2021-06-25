package validation

import (
	"../database"
	"fmt"
	"net/mail"
)

func UserExist(username string) bool {
	// return false when user exit, true when he doesn't exist
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

func valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func EmailExist(email string) bool {
	// return false when email exit, true when he doesn't exist
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

	if !valid(email) {
		return false
	}

	if !rows.Next() {
		fmt.Println("user with this email doesn't exit, you can create it")
		return true
	}
	return false
}
