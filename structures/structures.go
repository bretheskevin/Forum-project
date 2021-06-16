package structures

type Post struct {
	Title       string
	Content     string
	PublisherID int
	Category    string
}

type User struct {
	Username          string
	Email             string
	Password          string
	ProfilePictureURL string
	IsAdmin           bool
}
