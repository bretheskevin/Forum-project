package models

type User struct {
	ID                int
	UserName          string
	Email             string
	Password          string
	ProfilePictureURL string
	IsAdmin           bool
}

type NewUser struct {
	ID                int
	UserName          string
	Email             string
	NewPassword       string
	OldPassword       string
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

type PostReceive struct {
	Title    string
	Content  string
	Category string
	Topic    string
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
