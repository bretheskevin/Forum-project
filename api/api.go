package api

import (
	"../auth"
	"../models"
	"../utils/database"
	"../utils/password"
	"../utils/validation"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/mail"
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
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			return
		}

	} else {
		res := "Wrong password or email"
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			return
		}
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
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			return
		}
	}
	if token == "name" {
		res := "username already taken"
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			return
		}
	}

	if register == true {
		res := "user log"
		cookie := http.Cookie{Name: "token", Value: token, Expires: time.Now().Add(time.Minute * 60)}
		http.SetCookie(w, &cookie)
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			return
		}

	}
}

func getPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, _ := vars["id"]
	id, _ := strconv.Atoi(key)
	post, _ := database.GetPost(id)
	jsonFormat, _ := json.Marshal(post)

	w.Header().Set("content-type", "application/json")
	_, err := w.Write(jsonFormat)
	if err != nil {
		return
	}

}

func getPosts(w http.ResponseWriter, r *http.Request) {
	posts := database.GetPosts()
	res, _ := json.Marshal(posts)
	w.Header().Set("content-type", "application/json")
	_, err := w.Write(res)
	if err != nil {
		return
	}
}

func getPostsByCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	category, _ := vars["category"]
	posts := database.GetPostsByCategory(category)
	jsonFormat, _ := json.Marshal(posts)
	w.Header().Set("content-type", "application/json")
	_, err := w.Write(jsonFormat)
	if err != nil {
		return
	}
}

func getPostsByPublisher(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	publisherID, _ := strconv.Atoi(vars["id"])
	posts := database.GetPostsByPublisher(publisherID)
	jsonFormat, _ := json.Marshal(posts)
	w.Header().Set("content-type", "application/json")
	_, err := w.Write(jsonFormat)
	if err != nil {
		return
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	users := database.GetUsers()
	res, _ := json.Marshal(users)
	w.Header().Set("content-type", "application/json")
	_, err := w.Write(res)
	if err != nil {
		return
	}
}

func getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, _ := vars["id"]
	id, _ := strconv.Atoi(key)
	post, _ := database.GetUser(id)
	jsonFormat, _ := json.Marshal(post)

	w.Header().Set("content-type", "application/json")
	_, err := w.Write(jsonFormat)
	if err != nil {
		return
	}
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
	_, err := w.Write(res)
	if err != nil {
		return
	}
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
	_, err := w.Write(res)
	if err != nil {
		return
	}
}

func createPost(w http.ResponseWriter, r *http.Request) {
	jwtToken := getJWT(w, r)
	body, err := ioutil.ReadAll(r.Body)
	checkErr(err)
	var formattedBody models.PostReceive
	err = json.Unmarshal(body, &formattedBody)
	checkErr(err)
	category := formattedBody.Category + "/" + formattedBody.Topic
	database.AddPost(models.Post{
		Title:       formattedBody.Title,
		Content:     formattedBody.Content,
		Category:    category,
		PublisherID: int(jwtToken.(float64)),
	})
	er := json.NewEncoder(w).Encode("post create")
	if er != nil {
		return
	}
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

func updateUser(w http.ResponseWriter, r *http.Request) {
	userID := getJWT(w, r)
	body, err := ioutil.ReadAll(r.Body)
	checkErr(err)
	var formattedBody models.NewUser
	err = json.Unmarshal(body, &formattedBody)
	checkErr(err)

	// check if updater is owner
	user, exist := database.GetUser(formattedBody.ID)
	if exist {
		if user.ID != int(userID.(float64)) {
			err := json.NewEncoder(w).Encode("you are not the owner of this account")
			if err != nil {
				return
			}
			return
		}
	}

	// check if email already exist
	email := validation.EmailExist(formattedBody.Email)
	if !email {
		err := json.NewEncoder(w).Encode("email is already taken")
		if err != nil {
			return
		}
	}
	//check if email is valid
	_, emailErr := mail.ParseAddress(formattedBody.Email)
	if emailErr != nil {
		err := json.NewEncoder(w).Encode("email is not valid")
		if err != nil {
			return
		}
		return
	}

	username := validation.UserExist(formattedBody.UserName)
	if !username {
		err := json.NewEncoder(w).Encode("username is already taken")
		if err != nil {
			return
		}
	}
	if len(formattedBody.UserName) < 5 {
		err := json.NewEncoder(w).Encode("username is to short")
		if err != nil {
			return
		}
	}

	oldPass := formattedBody.OldPassword
	err = password.CheckPassword(oldPass, user.Password)

	if err != nil {
		err := json.NewEncoder(w).Encode("wrong password")
		if err != nil {
			return
		}
		return
	}

	newPassword, err := password.HashPassword(formattedBody.NewPassword)
	if err != nil {
		return
	}
	updateUser := database.UpdateUser(formattedBody.ID, formattedBody.UserName, formattedBody.Email, newPassword, formattedBody.ProfilePictureURL)

	if updateUser {
		err := json.NewEncoder(w).Encode("user updated")
		if err != nil {
			return
		}
	} else {
		err := json.NewEncoder(w).Encode("user not updated")
		if err != nil {
			return
		}
	}
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	publisherID := getJWT(w, r)
	body, err := ioutil.ReadAll(r.Body)
	checkErr(err)
	var formattedBody models.Post
	err = json.Unmarshal(body, &formattedBody)
	checkErr(err)

	// check if updater is owner of the post
	post, exist := database.GetPost(formattedBody.ID)
	if exist {
		if post.PublisherID != int(publisherID.(float64)) {
			json.NewEncoder(w).Encode("you are not the owner of this post")
			return
		}
	}

	// updatePost
	postUpdated := database.UpdatePost(formattedBody.ID, formattedBody.Title, formattedBody.Content, formattedBody.Category)
	if postUpdated {
		json.NewEncoder(w).Encode("post updated")
	} else {
		json.NewEncoder(w).Encode("post not updated")
	}
}

func Start(router *mux.Router) {
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/post", updatePost).Methods("PATCH")
	router.HandleFunc("/user", updateUser).Methods("PATCH")
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
