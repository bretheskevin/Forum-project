package api

import (
	"../auth"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
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

func login(w http.ResponseWriter, r *http.Request) {
	/*
		if err := r.ParseForm(); err != nil {
			fmt.Println(err)
			return
		}
		email := r.FormValue("email")
		pass := r.FormValue("password")
	*/
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

func Start(router *mux.Router) {
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/register", register).Methods("POST")
}

/*
body, err := ioutil.ReadAll(r.Body)
if err != nil {
fmt.Println(err)
}

var formattedBody Login
err = json.Unmarshal(body, &formattedBody)
if err != nil {
fmt.Println(err)
}*/
