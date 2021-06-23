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
		"is_admin"	BOOLEAN NOT NULL,
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

func GetPost(ID int) (models.Post, bool) {
	db := Connect()
	post := models.Post{}
	defer db.Close()
	stmt, err := db.Prepare("SELECT * FROM posts WHERE id=?")

	if err != nil {
		fmt.Println(err)
		return post, false
	}
	defer stmt.Close()

	rows, err := stmt.Query(ID)
	if err != nil {
		fmt.Println(err)
		return post, false
	}

	defer rows.Close()

	var id int
	var content string
	var title string
	var publisherID int
	var category string

	if rows.Next() {
		rows.Scan(&id, &title, &content, &publisherID, &category)
		post = models.Post{ID: id, Title: title, Content: content, PublisherID: publisherID, Category: category}
		return post, true
	}
	return post, false

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
		INSERT INTO users (username,email,password,profile_picture_url , is_admin)
		VALUES(?, ?, ?, ? ,?)
	`)
	if err != nil {
		log.Fatalln(err)
		return
	}
	stmt.Exec(user.UserName, user.Email, user.Password, "/images/default-pp.jpg", false)
	fmt.Println("User " + user.UserName + " added !")
}

func GetUser(ID int) (models.User, bool) {
	db := Connect()
	user := models.User{}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(db)
	stmt, err := db.Prepare("SELECT id,username,email,profile_picture_url,is_admin FROM users WHERE id=?")

	if err != nil {
		fmt.Println(err)
		return user, false
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(stmt)

	rows, err := stmt.Query(ID)
	if err != nil {
		fmt.Println(err)
		return user, false
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(rows)

	var id int
	var username string
	var email string
	var profilePictureUrl string
	var isAdmin bool

	if rows.Next() {
		err := rows.Scan(&id, &username, &email, &profilePictureUrl, &isAdmin)
		if err != nil {
			return models.User{}, false
		}
		user = models.User{
			ID:                id,
			UserName:          username,
			Email:             email,
			Password:          "secret",
			ProfilePictureURL: profilePictureUrl,
			IsAdmin:           isAdmin,
		}
		return user, true
	}
	return user, false
}

func GetUsers() []models.User {
	db := Connect()
	var users []models.User
	rows, _ := db.Query(
		`SELECT id,username,email,profile_picture_url,is_admin FROM users`)
	var id int
	var username string
	var email string
	var profilePictureUrl string
	var isAdmin bool

	for rows.Next() {
		err := rows.Scan(&id, &username, &email, &profilePictureUrl, &isAdmin)
		if err != nil {
			return nil
		}
		users = append(users, models.User{
			ID:                id,
			UserName:          username,
			Email:             email,
			Password:          "secret",
			ProfilePictureURL: profilePictureUrl,
			IsAdmin:           isAdmin,
		})
	}
	return users
}

func GetPostsByCategory(Category string) []models.Post {
	db := Connect()
	posts := []models.Post{}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(db)
	stmt, err := db.Prepare("SELECT * FROM posts WHERE category=?")

	rows, err := stmt.Query(Category)
	if err != nil {
		fmt.Println(err)
		return posts
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(rows)

	var id int
	var content string
	var title string
	var publisherID int
	var category string

	for rows.Next() {
		err := rows.Scan(&id, &title, &content, &publisherID, &category)
		if err != nil {
			return nil
		}
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

func GetPostsByPublisher(PublisherID int) []models.Post {
	db := Connect()
	posts := []models.Post{}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(db)
	stmt, err := db.Prepare("SELECT * FROM posts WHERE publisher_id=?")

	rows, err := stmt.Query(PublisherID)
	if err != nil {
		fmt.Println(err)
		return posts
	}

	defer rows.Close()

	var id int
	var content string
	var title string
	var publisherID int
	var category string

	for rows.Next() {
		err := rows.Scan(&id, &title, &content, &publisherID, &category)
		if err != nil {
			return nil
		}
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
