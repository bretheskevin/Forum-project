package database

import (
	"../../models"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func Connect() *sql.DB {
	db, err := sql.Open("sqlite3", "utils/database/data.db")
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

func UserTable(db *sql.DB) {
	stmt, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS "users" (
		"id"	INTEGER NOT NULL UNIQUE,
		"username"	TEXT NOT NULL UNIQUE,
		"email"	TEXT NOT NULL UNIQUE,
		"password"	INTEGER NOT NULL,
		"profile_picture_url"	TEXT NOT NULL,
		"admin"	BOOLEAN NOT NULL,
		PRIMARY KEY("id" AUTOINCREMENT)
		);
	`)
	if err != nil {
		log.Fatalln(err)
	}
	stmt.Exec()
}

func PostTable(db *sql.DB) {
	stmt, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS "posts" (
			"id"	INTEGER NOT NULL UNIQUE,
			"title"	TEXT NOT NULL,
			"content"	TEXT NOT NULL,
			"publisher_id"	INTEGER NOT NULL,
			"category"	TEXT NOT NULL,
			PRIMARY KEY("ID" AUTOINCREMENT)
		);
	`)
	if err != nil {
		log.Fatalln(err)
	}

	stmt.Exec()
}

func GetPosts() []models.Post {
	db := Connect()
	posts := []models.Post{}
	rows, _ := db.Query(
		`SELECT * FROM posts`)
	var id int
	var content string
	var title string
	var publisherID int
	var category string

	for rows.Next() {
		rows.Scan(&id, &title, &content, &publisherID, &category)
		posts = append(posts, models.Post{
			ID:          id,
			Title:       title,
			Content:     content,
			PublisherID: publisherID,
			Category:    category,
		})
	}
	return posts
}

func InitTable(db *sql.DB) {
	UserTable(db)
	PostTable(db)
}

func AddPost(database *sql.DB, post models.Post) {
	stmt, err := database.Prepare(`
		INSERT INTO posts (title, content, publisher_id, category)
		VALUES(?, ?, ?, ?)
	`)
	if err != nil {
		log.Fatalln(err)
		return
	}
	stmt.Exec(post.Title, post.Content, post.PublisherID, post.Category)
	fmt.Println("Post " + post.Title + " added !")
}

func AddUser(database *sql.DB, user models.User) {
	stmt, err := database.Prepare(`
		INSERT INTO users (username,email,password, profile_picture_url, admin)
		VALUES(?, ?, ?, ?)
	`)
	if err != nil {
		log.Fatalln(err)
		return
	}
	default_pp := "/"
	stmt.Exec(user.UserName, user.Email, user.Password, default_pp, false)
	fmt.Println("User " + user.UserName + " added !")
}
