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

func Init(database *sql.DB) {
	statement, statementErr := database.Prepare(`
		CREATE TABLE IF NOT EXISTS "postsfeed" (
			"ID"	INTEGER NOT NULL UNIQUE,
			"title"	TEXT NOT NULL,
			"content"	TEXT NOT NULL,
			"publisherID"	INTEGER NOT NULL,
			"publisherName"	TEXT NOT NULL,
			PRIMARY KEY("ID" AUTOINCREMENT)
		);
	`)
	if statementErr != nil {
		log.Fatalln(statementErr)
		return
	}

	statement.Exec()
}

func AddPost(database *sql.DB, user structures.User, post structures.Post) {
	statement, statementErr := database.Prepare(`
		INSERT INTO postsfeed (title, content, publisherID, publisherName)
		VALUES(?, ?, ?, ?)
	`)
	if statementErr != nil {
		log.Fatalln(statementErr)
		return
	}

	statement.Exec(post.Title, post.Content, user.ID, user.Username)

	fmt.Println("test")
}
