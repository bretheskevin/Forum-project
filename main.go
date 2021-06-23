package main

import (
	"./api"
	"./routes"
	"./utils/database"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	port := ":8080"

	// connect to the db
	db := database.Connect()
	database.InitTable(db)

	router := mux.NewRouter()
	api.Start(router)
	routes.Start(router)

	// test
	//test.CreatePosts()
	//posts := database.GetPosts()
	//fmt.Println(posts)
	//post, exist := database.GetPost(10)
	//fmt.Println(post,exist)

	// allow the server to access to the files

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))
	//http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))
	//http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("public/images"))))
	//http.Handle("/scripts/", http.StripPrefix("/scripts/", http.FileServer(http.Dir("public/scripts"))))
	//http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("public/styles"))))

	//start the server and use fmt to print the errors
	fmt.Println("Listening on http://localhost" + port)
	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatal(err)
		return
	}

}
