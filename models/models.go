package models

import "github.com/dgrijalva/jwt-go"

type User struct {
	ID                int
	UserName          string
	Email             string
	Password          string
	ProfilePictureURL string
	IsAdmin           bool
}

type Post struct {
	ID          int
	Title       string
	Content     string
	PublisherID int
	Category    string
}

type Login struct {
	Email    string
	Password string
}

type Register struct {
	Email    string
	Username string
	Password string
}

type Claims struct {
	Username string
	jwt.StandardClaims
}
