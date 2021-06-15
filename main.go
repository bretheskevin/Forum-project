package main

import (
	"./database"
	"./routes"
	"./structures"
	"fmt"
	"log"
	"net/http"
)

func main() {

	db := database.OpenDB("database/forum.db")
	database.Init(db) // create the different tables

	user := structures.User{
		ID:       1,
		Username: "Kevin",
	}

	post := structures.Post{
		Title:   "Les aventures",
		Content: "INCROYABLE !!!",
	}

	database.AddPost(db, user, post)

	return

	port := ":8080"

	// allow the server to access to the files
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("public/images"))))
	http.Handle("/scripts/", http.StripPrefix("/scripts/", http.FileServer(http.Dir("public/scripts"))))
	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("public/styles"))))

	//This method takes the URL path and a function that will show the page
	http.HandleFunc("/", routes.HomePage)
	http.HandleFunc("/homepage", routes.HomePage)
	http.HandleFunc("/feed", routes.FeedPage)
	http.HandleFunc("/admin/dashboard/", routes.DashboardPage)
	http.HandleFunc("/admin/dashboard/posts", routes.DashboardPostsPage)
	http.HandleFunc("/register", routes.RegisterPage)
	http.HandleFunc("/login", routes.LoginPage)

	//start the server and use fmt to print the errors
	fmt.Println("Listening on http://localhost" + port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
}
