package api

import (
	"../auth"
	"../utils/database"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type Login struct {
	Email    string
	Password string
}

type Register struct {
	Email    string
	Username string
	Password string
}

type Post struct {
	ID int
}

func login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	//r.Cookie("token")

	var formattedBody Login
	err = json.Unmarshal(body, &formattedBody)
	if err != nil {
		fmt.Println(err)
	}

	login, token := auth.Login(formattedBody.Email, formattedBody.Password)
	if login == true {
		res := "user log"
		cookie := http.Cookie{Name: "token", Value: token, Expires: time.Now().Add(time.Minute * 60)}
		http.SetCookie(w, &cookie)
		json.NewEncoder(w).Encode(res)

	} else {
		res := "Wrong password or email"
		json.NewEncoder(w).Encode(res)
	}
}

func register(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	//r.Cookie("token")

	var formattedBody Register
	err = json.Unmarshal(body, &formattedBody)
	if err != nil {
		fmt.Println(err)
	}
	register, token := auth.Register(formattedBody.Username, formattedBody.Email, formattedBody.Password)

	if token == "email" {
		res := "email already registered"
		json.NewEncoder(w).Encode(res)
	}
	if token == "name" {
		res := "username already taken"
		json.NewEncoder(w).Encode(res)
	}

	if register == true {
		res := "user log"
		cookie := http.Cookie{Name: "token", Value: token, Expires: time.Now().Add(time.Minute * 60)}
		http.SetCookie(w, &cookie)
		json.NewEncoder(w).Encode(res)

	}
}

func getPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, _ := vars["id"]
	id, _ := strconv.Atoi(key)
	post, _ := database.GetPost(id)
	jsonFormat, _ := json.Marshal(post)

	w.Header().Set("content-type", "application/json")
	w.Write(jsonFormat)

}

func getPosts(w http.ResponseWriter, r *http.Request) {
	posts := database.GetPosts()
	res, _ := json.Marshal(posts)
	w.Header().Set("content-type", "application/json")
	w.Write(res)
}

func getPostsByCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	category, _ := vars["category"]
	posts := database.GetPostsByCategory(category)
	jsonFormat, _ := json.Marshal(posts)
	w.Header().Set("content-type", "application/json")
	w.Write(jsonFormat)
}

func getPostsByPublisher(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	publisherID, _ := strconv.Atoi(vars["id"])
	posts := database.GetPostsByPublisher(publisherID)
	jsonFormat, _ := json.Marshal(posts)
	w.Header().Set("content-type", "application/json")
	w.Write(jsonFormat)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	users := database.GetUsers()
	res, _ := json.Marshal(users)
	w.Header().Set("content-type", "application/json")
	w.Write(res)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, _ := vars["id"]
	id, _ := strconv.Atoi(key)
	post, _ := database.GetUser(id)
	jsonFormat, _ := json.Marshal(post)

	w.Header().Set("content-type", "application/json")
	w.Write(jsonFormat)
}

func Start(router *mux.Router) {
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/register", register).Methods("POST")
	router.HandleFunc("/post/{id}", getPost).Methods("GET")
	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/user/{id}", getUser).Methods("GET")
	router.HandleFunc("/posts/{category}", getPostsByCategory).Methods("GET")
	router.HandleFunc("/posts/user/{id}", getPostsByPublisher).Methods("GET")
}
