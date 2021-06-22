package main

import (
	"./api"
	"./routes"
	"./utils/database"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// connect to the db
	db := database.Connect()
	database.InitTable(db)

	api.Start()

	// test create post
	//post := models.Post{Title: "FirstArticle",Content: "Hello",PublisherID: 19,Category: "test"}
	//database.AddPost(db, post)
	//posts := database.GetPosts()
	//fmt.Println(posts)
	//database.AddPost(models.Post{Title: "SecondArticle",Content: "Hello2"})

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
