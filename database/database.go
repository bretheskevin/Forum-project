package database

import (
	"../structures"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func OpenDB(path string) *sql.DB {
	db, databaseErr := sql.Open("sqlite3", path)
	if databaseErr != nil {
		log.Fatalln(databaseErr)
	}
	return db
}

func createPostsFeedTable(database *sql.DB) {
	// posts feed
	statement, statementErr := database.Prepare(`
		CREATE TABLE IF NOT EXISTS "postsfeed" (
			"id"	INTEGER NOT NULL UNIQUE,
			"title"	TEXT NOT NULL,
			"content"	TEXT NOT NULL,
			"publisher_id"	INTEGER NOT NULL,
			"category"	TEXT NOT NULL,
			PRIMARY KEY("ID" AUTOINCREMENT)
		);
	`)
	if statementErr != nil {
		log.Fatalln(statementErr)
		return
	}

	statement.Exec()
}

func createUsersTable(database *sql.DB) {
	statement, statementErr := database.Prepare(`
		CREATE TABLE IF NOT EXISTS "users" (
		"id"	INTEGER NOT NULL UNIQUE,
		"username"	TEXT NOT NULL,
		"email"	TEXT NOT NULL,
		"password"	INTEGER NOT NULL,
		"profile_picture"	INTEGER NOT NULL,
		"is_admin" BOOLEAN NOT NULL,
		PRIMARY KEY("ID" AUTOINCREMENT)
		);
	`)
	if statementErr != nil {
		log.Fatalln(statementErr)
		return
	}

	statement.Exec()
}

func Init(database *sql.DB) {
	createPostsFeedTable(database)
	createUsersTable(database)
}

func AddPost(database *sql.DB, post structures.Post) {
	statement, statementErr := database.Prepare(`
		INSERT INTO postsfeed (title, content, publisher_id, category)
		VALUES(?, ?, ?, ?)
	`)
	if statementErr != nil {
		log.Fatalln(statementErr)
		return
	}

	statement.Exec(post.Title, post.Content, post.PublisherID, post.Category)

	fmt.Println("Post \"" + post.Title + "\" added !")

}

func AddUser(database *sql.DB, user structures.User) {
	statement, statementErr := database.Prepare(`
		INSERT INTO users (username, email, password, profile_picture, is_admin)
		VALUES(?, ?, ?, ?, ?)
	`)
	if statementErr != nil {
		log.Fatalln(statementErr)
		return
	}

	statement.Exec(user.Username, user.Email, user.Password, user.ProfilePictureURL, user.IsAdmin)

	fmt.Println("User \"" + user.Username + "\" added !")
}
