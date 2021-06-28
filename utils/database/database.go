package database

import (
	"../../models"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Connect() *sql.DB {
	// connect to the db
	db, err := sql.Open("sqlite3", "utils/database/data.db")
	checkErr(err)
	return db
}

func UserTable(db *sql.DB) {
	// create User table in sqlite
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
	checkErr(err)
	stmt.Exec()
}

func PostTable(db *sql.DB) {
	// create Post table in sqlite
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
	checkErr(err)

	stmt.Exec()
}

func GetPosts() []models.Post {
	// get all posts, return an array of Posts
	db := Connect()
	posts := []models.Post{} // init array
	rows, _ := db.Query(
		`SELECT * FROM posts`) // select all from posts
	var id int
	var content string
	var title string
	var publisherID int
	var category string

	for rows.Next() { // for each posts founded
		rows.Scan(&id, &title, &content, &publisherID, &category) // scan all element in row
		posts = append(posts, models.Post{                        // append result in posts array
			ID:          id,
			Title:       title,
			Content:     content,
			PublisherID: publisherID,
			Category:    category,
		})
	}
	return posts // return posts array
}

func GetPost(ID int) (models.Post, bool) {
	/*
		get Post by PostID, return post, and bool,
		if post exist, return true, else return false
	*/

	db := Connect()
	post := models.Post{} // init post
	defer db.Close()
	stmt, err := db.Prepare("SELECT * FROM posts WHERE id=?") // select all form post with params ID

	if err != nil {
		fmt.Println(err)
		return post, false
	}
	defer stmt.Close()

	rows, err := stmt.Query(ID) // select from ID
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

	if rows.Next() { // if post founded, scan and add info in post
		rows.Scan(&id, &title, &content, &publisherID, &category)
		post = models.Post{ID: id, Title: title, Content: content, PublisherID: publisherID, Category: category}
		return post, true // return post and true
	}
	return post, false // doesn't found any post with this id, return empty post, and false

}

func InitTable(db *sql.DB) {
	// init database Table
	UserTable(db)
	PostTable(db)
}

func AddPost(post models.Post) {
	// add post into the Database, take a post in params
	db := Connect()
	stmt, err := db.Prepare(`
		INSERT INTO posts (title, content, publisher_id, category)
		VALUES(?, ?, ?, ?)
	`) // prepare the database
	if err != nil {
		log.Fatalln(err)
		return
	}
	stmt.Exec(post.Title, post.Content, post.PublisherID, post.Category) // add post in the database
	fmt.Println("Post " + post.Title + " added !")
}

func AddUser(user models.User) {
	// add user into the Database, take a user in params
	db := Connect()
	stmt, err := db.Prepare(`
		INSERT INTO users (username,email,password,profile_picture_url , is_admin)
		VALUES(?, ?, ?, ? ,?)
	`) // prepare the database
	if err != nil {
		log.Fatalln(err)
		return
	}
	stmt.Exec(user.UserName, user.Email, user.Password, "/images/default-pp.jpg", false) // add user in the database
	fmt.Println("User " + user.UserName + " added !")
}

func GetUser(ID int) (models.User, bool) {
	/*
		get user by ID, return an user and bool
		true if user exist, else false
	*/
	db := Connect()
	user := models.User{} // init user

	defer func(db *sql.DB) { // defer close database
		err := db.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(db)
	stmt, err := db.Prepare("SELECT id,username,email,profile_picture_url,is_admin FROM users WHERE id=?") // prepare database

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

	rows, err := stmt.Query(ID) // found rows with ID
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

	if rows.Next() { // if a row is founded
		err := rows.Scan(&id, &username, &email, &profilePictureUrl, &isAdmin) // scan user info
		if err != nil {
			return models.User{}, false
		}
		user = models.User{ // append info in user result
			ID:                id,
			UserName:          username,
			Email:             email,
			Password:          "secret",
			ProfilePictureURL: profilePictureUrl,
			IsAdmin:           isAdmin,
		}
		return user, true // return user and true
	}
	return user, false
}

func GetUsers() []models.User {
	/*
		getUsers function return an array of users
	*/
	db := Connect()
	var users []models.User // init users array
	rows, _ := db.Query(
		`SELECT id,username,email,profile_picture_url,is_admin FROM users`) // select all and stock in rows
	var id int
	var username string
	var email string
	var profilePictureUrl string
	var isAdmin bool

	for rows.Next() { // for each rows,
		err := rows.Scan(&id, &username, &email, &profilePictureUrl, &isAdmin) // scan info
		if err != nil {
			return nil
		}
		users = append(users, models.User{ // append info in users array
			ID:                id,
			UserName:          username,
			Email:             email,
			Password:          "secret",
			ProfilePictureURL: profilePictureUrl,
			IsAdmin:           isAdmin,
		})
	}
	return users // return users array
}

func GetPostsByCategory(Category string) []models.Post {
	/*
		getPostByCategory take Category in params,
		and return list of post who patch with this input category
	*/

	db := Connect()
	posts := []models.Post{} // init posts array
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(db)
	stmt, err := db.Prepare("SELECT * FROM posts WHERE category=?") // prepare satement where category match with in input

	rows, err := stmt.Query(Category) // stock matching result in rows
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

	for rows.Next() { // for each founded rows
		err := rows.Scan(&id, &title, &content, &publisherID, &category) // scan post info in row
		if err != nil {
			return nil
		}
		posts = append(posts, models.Post{ // append post into posts array
			ID:          id,
			Title:       title,
			Content:     content,
			PublisherID: publisherID,
			Category:    category,
		})
	}
	return posts // return posts array
}

func GetPostsByPublisher(PublisherID int) []models.Post {
	/*
		GetPostsByPublisher take publisherId in params,
		and return list of Posts wrote by the publisherID
	*/
	db := Connect()
	posts := []models.Post{} // init posts array
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(db)
	stmt, err := db.Prepare("SELECT * FROM posts WHERE publisher_id=?") // prepare stmt

	rows, err := stmt.Query(PublisherID) // stock match rows where publisherId is equal to input, and stock result in rows
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

	for rows.Next() { // for each rows
		err := rows.Scan(&id, &title, &content, &publisherID, &category) // scan post info
		if err != nil {
			return nil
		}
		posts = append(posts, models.Post{ // append post into Posts array
			ID:          id,
			Title:       title,
			Content:     content,
			PublisherID: publisherID,
			Category:    category,
		})
	}
	return posts // return posts array
}

func DeleteUser(ID int) bool {
	/*
		DeleteUser take id of user in params, and return true if user has been deleted
		if user has been deleted, delete all his posts
	*/
	db := Connect()
	stmt, err := db.Prepare("DELETE FROM users WHERE id = ?") // delete user where id is equal to ID
	checkErr(err)
	result, errExec := stmt.Exec(ID)
	checkErr(errExec)
	rowAffected, errRow := result.RowsAffected() // check if rows is affected
	checkErr(errRow)
	if rowAffected != 0 { // if rows is affected, user has been deleted, we can delete all of posts wrote by him
		DeletePostByPublisher(ID)
		return true
	}
	return false
}

func DeletePost(ID int) bool {
	// deletePost take post ID in params, and return true if post has been deleted
	db := Connect()
	stmt, err := db.Prepare("DELETE FROM posts WHERE id = ?") // stmt database
	checkErr(err)
	result, errExec := stmt.Exec(ID) // exec with id
	checkErr(errExec)
	rowAffected, errRow := result.RowsAffected() // check if rows is affected
	checkErr(errRow)
	if rowAffected != 0 { // if rows is affected, post as been deleted, and return true
		return true
	}
	return false
}

func DeletePostByPublisher(PublisherID int) bool {
	/*
		// deletePostByPublisher take post ID in params, and return true if post has been deleted
	*/
	db := Connect()
	stmt, err := db.Prepare("DELETE FROM posts WHERE publisher_id = ?") // prepare stmt
	checkErr(err)
	result, errExec := stmt.Exec(PublisherID) // exec with publisher id
	checkErr(errExec)
	rowAffected, errRow := result.RowsAffected() // check if rows is affected
	checkErr(errRow)
	if rowAffected != 0 { // if rows is affected, post as been deleted, and return true
		return true
	}
	return false
}

func GetUserByEmail(Email string) (models.User, bool) {
	db := Connect()
	user := models.User{} // init user
	defer db.Close()
	stmt, err := db.Prepare("SELECT id,username,email FROM users WHERE email=?")

	if err != nil {
		fmt.Println(err)
		return user, false
	}
	defer stmt.Close()

	rows, err := stmt.Query(Email)
	if err != nil {
		fmt.Println(err)
		return user, false
	}

	defer rows.Close()

	var id int
	var username string
	var email string

	if rows.Next() {
		rows.Scan(&id, &username, &email)
		user = models.User{ID: id, UserName: username, Email: email, Password: "secret"} // secret because we don't want to send hash password
		return user, true
	}
	return user, false
}

func UpdatePost(ID int, Title string, Content string, Category string) bool {
	db := Connect()
	stmt, err := db.Prepare("UPDATE posts SET title=?, content=?, category=? WHERE id=?") // prepare to update
	if err != nil {
		fmt.Println(err)
		return false
	}
	res, err := stmt.Exec(Title, Content, Category, ID) // exec with params
	if err != nil {
		fmt.Println(err)
		return false
	}
	a, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return false
	}
	if a != 0 { // check id rows is affected, if yes, return true, post has been updated
		return true
	}
	return false
}

//////// Not fonctionnale ////////////
func UpdateUser(ID int, Username string, Email string, Password string, Image string) bool {
	db := Connect()
	sqlStatement := `
UPDATE users
SET username = $2, email = $3, password = $4, profile_picture_url = $5
WHERE id = $1;`
	defer db.Close()
	_, err := db.Exec(sqlStatement, ID, Username, Email, Password, Image)
	if err != nil {
		panic(err)
		return false
	}
	return true
}

//////////////////////////////////////
