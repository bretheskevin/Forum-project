package routes

import (
	"html/template"
	"log"
	"net/http"
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
