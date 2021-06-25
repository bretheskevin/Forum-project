package api

import (
	"../auth"
	"../models"
	"../utils/database"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func checkErr(err error) {
	// short check error function
	if err != nil {
		panic(err)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body) // read body
	checkErr(err)

	var formattedBody models.Login // formatted body with Login models

	err = json.Unmarshal(body, &formattedBody) // parse json
	checkErr(err)

	login, token := auth.Login(formattedBody.Email, formattedBody.Password) // try to login

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
	body, err := ioutil.ReadAll(r.Body) // read body
	checkErr(err)

	var formattedBody models.Register
	err = json.Unmarshal(body, &formattedBody) // parse json
	checkErr(err)
	register, token := auth.Register(formattedBody.Username, formattedBody.Email, formattedBody.Password) //try to register

	// check email error
	if token == "email" {
		res := "email already registered"
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			return
		}
	}
	// check username error
	if token == "name" {
		res := "username already taken"
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			return
		}
	}

	// if register is possible
	if register == true {
		res := "user log"
		// set cookie
		cookie := http.Cookie{Name: "token", Value: token, Expires: time.Now().Add(time.Minute * 60)}
		http.SetCookie(w, &cookie)
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			return
		}

	}
}

func getPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)                 // get var in url
	key, _ := vars["id"]                // take id var
	id, _ := strconv.Atoi(key)          // convert to int
	post, exist := database.GetPost(id) // get post with id
	if !exist {                         // check if he exist
		return
	}
	jsonFormat, _ := json.Marshal(post)

	w.Header().Set("content-type", "application/json")
	_, err := w.Write(jsonFormat) // print
	if err != nil {
		return
	}

}

func getPosts(w http.ResponseWriter, r *http.Request) {
	posts := database.GetPosts()  // call getPosts who return list of posts
	res, _ := json.Marshal(posts) // parse
	w.Header().Set("content-type", "application/json")
	_, err := w.Write(res)
	if err != nil {
		return
	}
}

func getPostsByCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)             // get var in url
	category, _ := vars["category"] // take category var
	posts := database.GetPostsByCategory(category)
	jsonFormat, _ := json.Marshal(posts) // format
	w.Header().Set("content-type", "application/json")
	_, err := w.Write(jsonFormat) // print
	if err != nil {
		return
	}
}

func getPostsByPublisher(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)                        // get var in url
	publisherID, _ := strconv.Atoi(vars["id"]) // take id var
	posts := database.GetPostsByPublisher(publisherID)
	jsonFormat, _ := json.Marshal(posts) // format
	w.Header().Set("content-type", "application/json")
	_, err := w.Write(jsonFormat) // print
	if err != nil {
		return
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	users := database.GetUsers()  // just call methode
	res, _ := json.Marshal(users) // format res
	w.Header().Set("content-type", "application/json")
	_, err := w.Write(res) // print
	if err != nil {
		return
	}
}

func getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)                 // get var in url
	key, _ := vars["id"]                // take id var
	id, _ := strconv.Atoi(key)          // convert to int
	post, _ := database.GetUser(id)     // get user with id
	jsonFormat, _ := json.Marshal(post) // format

	w.Header().Set("content-type", "application/json")
	_, err := w.Write(jsonFormat) // print
	if err != nil {
		return
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	var message string
	vars := mux.Vars(r)                 // get var in url
	id, _ := strconv.Atoi(vars["id"])   // take id var and convert it to int
	isDelete := database.DeleteUser(id) // call deleteUser with id
	if isDelete {                       // if he can be deleted
		message = "user with id " + strconv.Itoa(id) + " has been delete"
	} else {
		message = "we can't found user with id: " + strconv.Itoa(id)
	}
	res, _ := json.Marshal(message) // format
	w.Header().Set("content-type", "application/json")
	_, err := w.Write(res) // print
	if err != nil {
		return
	}
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	var message string
	vars := mux.Vars(r)                 // get var in url
	id, _ := strconv.Atoi(vars["id"])   // take id var and convert it to int
	isDelete := database.DeletePost(id) // call deletePost with id
	if isDelete {                       // if he can be deleted
		message = "post with id " + strconv.Itoa(id) + " has been delete"
	} else {
		message = "we can't found post with id: " + strconv.Itoa(id)
	}
	res, _ := json.Marshal(message)
	w.Header().Set("content-type", "application/json")
	_, err := w.Write(res) // print
	if err != nil {
		return
	}
}

func createPost(w http.ResponseWriter, r *http.Request) {
	jwtToken := getJWT(w, r)            // get Id of current User
	body, err := ioutil.ReadAll(r.Body) // read body
	checkErr(err)
	var formattedBody models.PostReceive       // format body
	err = json.Unmarshal(body, &formattedBody) // format
	checkErr(err)
	category := formattedBody.Category + "-" + formattedBody.Topic // format category
	database.AddPost(models.Post{                                  // Add Post
		Title:       formattedBody.Title,
		Content:     formattedBody.Content,
		Category:    category,
		PublisherID: int(jwtToken.(float64)),
	})
	er := json.NewEncoder(w).Encode("post create") // print
	if er != nil {
		return
	}
}

func getJWT(w http.ResponseWriter, r *http.Request) interface{} {
	cookie, err := r.Cookie("token") // take cookie who named "token"
	checkErr(err)
	tokenString := cookie.Value // take value of token
	claims := jwt.MapClaims{}
	jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("TokenPassword"), nil
	})
	// ... error handling

	for key, val := range claims { // ranges claims Token
		//fmt.Printf("Key: %v, value: %v\n", key, val)
		if key == "user_id" { // return user_id value
			return val
		}
	}
	return -1
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	publisherID := getJWT(w, r)         // get id of current user
	body, err := ioutil.ReadAll(r.Body) // read body
	checkErr(err)
	var formattedBody models.Post // format the body
	err = json.Unmarshal(body, &formattedBody)
	checkErr(err)

	// check if updater is owner of the post
	post, exist := database.GetPost(formattedBody.ID) // getPost with id
	if exist {
		if post.PublisherID != int(publisherID.(float64)) { // check if current user is a owner of the post
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
	// authentication
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/register", register).Methods("POST")

	/// POST
	// get by id, by user id, by category, get all
	router.HandleFunc("/post/{id}", getPost).Methods("GET")
	router.HandleFunc("/posts/{category}", getPostsByCategory).Methods("GET")
	router.HandleFunc("/posts/user/{id}", getPostsByPublisher).Methods("GET")
	router.HandleFunc("/posts", getPosts).Methods("GET")

	// create post
	router.HandleFunc("/post", createPost).Methods("POST")

	// delete post
	router.HandleFunc("/post/{id}", deletePost).Methods("DELETE")

	/// USER
	// get by id, get all
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/user/{id}", getUser).Methods("GET")

	// delete by id
	router.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")

	/// BUG
	// update user, not ready 	router.HandleFunc("/user", updateUser).Methods("PATCH")
	router.HandleFunc("/post", updatePost).Methods("PATCH")
}

/// BUG ///
/*
func updateUser(w http.ResponseWriter, r *http.Request) {
	userID := getJWT(w, r)
	body, err := ioutil.ReadAll(r.Body)
	checkErr(err)
	var formattedBody models.NewUser
	err = json.Unmarshal(body, &formattedBody)
	checkErr(err)
	db := database.Connect()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	stmt, err := db.Prepare("SELECT * FROM users WHERE id=?")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {

		}
	}(stmt)

	rows, err := stmt.Query(userID)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	if !rows.Next() {
		fmt.Println("user not found")
		return
	}

	user := models.User{}
	if err = rows.Scan(&user.ID, &user.UserName, &user.Email, &user.Password, &user.ProfilePictureURL, &user.IsAdmin); err != nil {
		fmt.Println(err)
		return
	}

	hash := user.Password //hashed pass in the db
	fmt.Println(hash)
	fmt.Println(formattedBody.OldPassword)
	fmt.Println(formattedBody.NewPassword)

	errPass := password.CheckPassword(formattedBody.OldPassword, hash)
	if errPass != nil {
		err := json.NewEncoder(w).Encode("Wrong Password")
		if err != nil {
			return
		}
		return
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
*/
