package main

import (
	"fmt"
	"log"
	"net/http"

	"./routes"
)

func main() {
	// db := database.OpenDB("database/forum.db")
	// database.Init(db) // create the different tables

	// post := structures.Post{
	// 	Title:       "Les aventures de Kévin",
	// 	Content:     "INCROYABLE !!!",
	// 	PublisherID: 2,
	// 	Category:    "Cybersecurity",
	// }
	// database.AddPost(db, post)

	// user := structures.User{
	// 	Username:          "Kévin",
	// 	Email:             "kevin.brethes@ynov.com",
	// 	Password:          "Pa$$W0rD",
	// 	ProfilePictureURL: "https://cdn.discordapp.com/attachments/508258795877564416/778390220650709002/20201018_192834.jpg",
	// 	IsAdmin:           false,
	// }
	// database.AddUser(db, user)

	// return

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
	http.HandleFunc("/dashboard/", routes.DashboardPage)
	http.HandleFunc("/dashboard/posts", routes.DashboardPostsPage)
	http.HandleFunc("/register", routes.RegisterPage)
	http.HandleFunc("/login", routes.LoginPage)
	http.HandleFunc("/create-new-topic", routes.CreateTopic)

	//start the server and use fmt to print the errors
	fmt.Println("Listening on http://localhost" + port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
}
