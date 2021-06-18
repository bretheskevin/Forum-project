package main

import (
	"./api/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"strconv"
)

var posts []models.Post

func GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, i := range posts {
		if i.ID == params["id"] {
			json.NewEncoder(w).Encode(i)
			return
		}
	}
	json.NewEncoder(w).Encode(&models.Post{})
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post models.Post
	_ = json.NewDecoder(r.Body).Decode(&post)
	post.ID = strconv.Itoa(rand.Intn(1000000))
	posts = append(posts, post)
	json.NewEncoder(w).Encode(post)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range posts {
		if item.ID == params["id"] {
			posts = append(posts[:index], posts[index+1:]...)
			var post models.Post
			_ = json.NewDecoder(r.Body).Decode(&post)
			post.ID = params["id"]
			posts = append(posts, post)
			json.NewEncoder(w).Encode(post)
			return
		}
	}
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range posts {
		if item.ID == params["id"] {
			posts = append(posts[:index], posts[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(posts)
}

func main() {
	// create a new router
	r := mux.NewRouter()
	posts = append(posts, models.Post{ID: "1", Title: "Test", Content: "Hello World", Author: &models.User{
		Name: "gab", UUID: "1114", Email: "test@gmail.com",
	}})
	posts = append(posts, models.Post{ID: "2", Title: "Test2", Content: "Hello World2", Author: &models.User{
		Name: "gab2", UUID: "7819", Email: "test2@gmail.com",
	}})
	// create different route
	r.HandleFunc("/api/v1/posts", GetPosts).Methods("GET")
	r.HandleFunc("/api/v1/posts/{id}", GetPost).Methods("GET")
	r.HandleFunc("/api/v1/posts", CreatePost).Methods("POST")
	r.HandleFunc("/api/v1/posts{id}", UpdatePost).Methods("PUT")
	r.HandleFunc("/api/v1/posts{id}", DeletePost).Methods("DELETE")

	http.ListenAndServe(":8080", r)

}
