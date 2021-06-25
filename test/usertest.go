package test

import (
	"../models"
	"../utils/database"
	"../utils/password"
)

func CreateAccount() {
	users := []models.User{
		{UserName: "Gabriel", Email: "gabriel@cancel.fr"},
		{UserName: "Jean", Email: "jean@leroi.fr"},
		{UserName: "Paul", Email: "paul@lil.fr"},
		{UserName: "Boris", Email: "boris@poa.fr"},
		{UserName: "Steven", Email: "steven@nial.fr"},
	}

	for i := 0; i < len(users); i++ {
		newPass, _ := password.HashPassword(users[i].UserName)
		user := models.User{UserName: users[i].UserName, Email: users[i].Email, Password: newPass}
		database.AddUser(user)
	}
}
