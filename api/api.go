package api

import (
	"../auth"
	"../models"
	"../utils/database"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	fmt.Println(body)
	checkErr(err)
	//r.Cookie("token")

	var formattedBody models.Login
	err = json.Unmarshal(body, &formattedBody)
	checkErr(err)

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
	checkErr(err)
	//r.Cookie("token")

	var formattedBody models.Register
	err = json.Unmarshal(body, &formattedBody)
	checkErr(err)
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

func deleteUser(w http.ResponseWriter, r *http.Request) {
	var message string
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	isDelete := database.DeleteUser(id)
	if isDelete {
		message = "user with id " + strconv.Itoa(id) + " has been delete"
	} else {
		message = "we can't found user with id: " + strconv.Itoa(id)
	}
	res, _ := json.Marshal(message)
	w.Header().Set("content-type", "application/json")
	w.Write(res)
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	var message string
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	isDelete := database.DeletePost(id)
	if isDelete {
		message = "post with id " + strconv.Itoa(id) + " has been delete"
	} else {
		message = "we can't found post with id: " + strconv.Itoa(id)
	}
	res, _ := json.Marshal(message)
	w.Header().Set("content-type", "application/json")
	w.Write(res)
}

// create post function to create post in db

func createPost(w http.ResponseWriter, r *http.Request) {
	jwt := getJWT(w, r)
	db := database.Connect()
	body, err := ioutil.ReadAll(r.Body)
	checkErr(err)
	var formattedBody models.PostReceive
	err = json.Unmarshal(body, &formattedBody)
	checkErr(err)
	category := formattedBody.Category + "/" + formattedBody.Topic
	database.AddPost(db, models.Post{
		Title:       formattedBody.Title,
		Content:     formattedBody.Content,
		Category:    category,
		PublisherID: int(jwt.(float64)),
	})
	json.NewEncoder(w).Encode("post create")
}

func getJWT(w http.ResponseWriter, r *http.Request) interface{} {
	cookie, err := r.Cookie("token")
	checkErr(err)
	tokenString := cookie.Value
	claims := jwt.MapClaims{}
	jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("TokenPassword"), nil
	})
	// ... error handling

	// do something with decoded claims
	//fmt.Println("Token: ", token)
	for key, val := range claims {
		//fmt.Printf("Key: %v, value: %v\n", key, val)
		if key == "user_id" {
			return val
		}
	}
	return -1
}

// update user

// update post

func Start(router *mux.Router) {
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/register", register).Methods("POST")
	router.HandleFunc("/post/{id}", getPost).Methods("GET")
	router.HandleFunc("/post/{id}", deletePost).Methods("DELETE")
	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/user/{id}", getUser).Methods("GET")
	router.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")
	router.HandleFunc("/posts/{category}", getPostsByCategory).Methods("GET")
	router.HandleFunc("/posts/user/{id}", getPostsByPublisher).Methods("GET")
	router.HandleFunc("/post", createPost).Methods("POST")
}
