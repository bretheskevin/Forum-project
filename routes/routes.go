package routes

import (
	"../auth"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" && r.URL.Path != "/homepage" {
		errorHandler(w, r)
		return
	}

	page, err := template.ParseFiles("public/homepage.html")
	if err != nil {
		log.Fatal(err)
	}
	errorExecuteTemplate := page.ExecuteTemplate(w, "homepage.html", "")
	if errorExecuteTemplate != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func FeedPage(w http.ResponseWriter, r *http.Request) {
	page, err := template.ParseFiles("public/feed.html")
	if err != nil {
		log.Fatal(err)
	}
	errorExecuteTemplate := page.ExecuteTemplate(w, "feed.html", "")
	if errorExecuteTemplate != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func DashboardPage(w http.ResponseWriter, r *http.Request) {
	page, err := template.ParseFiles("public/dashboard.html")
	if err != nil {
		log.Fatal(err)
	}
	errorExecuteTemplate := page.ExecuteTemplate(w, "dashboard.html", "")
	if errorExecuteTemplate != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func DashboardPostsPage(w http.ResponseWriter, r *http.Request) {
	page, err := template.ParseFiles("public/dashboard-posts.html")
	if err != nil {
		log.Fatal(err)
	}
	errorExecuteTemplate := page.ExecuteTemplate(w, "dashboard-posts.html", "")
	if errorExecuteTemplate != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func RegisterPage(w http.ResponseWriter, r *http.Request) {
	page, err := template.ParseFiles("public/register.html")
	if err != nil {
		log.Fatal(err)
	}
	errorExecuteTemplate := page.ExecuteTemplate(w, "register.html", "")
	if errorExecuteTemplate != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func LoginPage(w http.ResponseWriter, r *http.Request) {

	page, err := template.ParseFiles("public/login.html")
	if err != nil {
		log.Fatal(err)
	}
	errorExecuteTemplate := page.ExecuteTemplate(w, "login.html", "")
	if errorExecuteTemplate != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := r.ParseForm(); err != nil {
		fmt.Println(err)
		return
	}

	email := r.FormValue("email")
	pass := r.FormValue("password")

	login, token := auth.Login(email, pass)
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

func CreateTopic(w http.ResponseWriter, r *http.Request) {
	page, err := template.ParseFiles("public/new-topic.html")
	if err != nil {
		log.Fatal(err)
	}
	errorExecuteTemplate := page.ExecuteTemplate(w, "new-topic.html", "")
	if errorExecuteTemplate != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)

	page, err := template.ParseFiles("public/404.html")
	if err != nil {
		log.Fatal(err)
	}
	errorExecuteTemplate := page.ExecuteTemplate(w, "404.html", "")
	if errorExecuteTemplate != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
