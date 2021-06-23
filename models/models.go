package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
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

type ResponseUser struct {
	ID       int
	UserName string
	Email    string
}
