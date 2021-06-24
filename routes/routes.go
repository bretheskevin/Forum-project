package routes

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
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

func feedPage(w http.ResponseWriter, r *http.Request) {
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

func dashboardPage(w http.ResponseWriter, r *http.Request) {
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

func dashboardPostsPage(w http.ResponseWriter, r *http.Request) {
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

func registerPage(w http.ResponseWriter, r *http.Request) {
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

func loginPage(w http.ResponseWriter, r *http.Request) {

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

func createTopic(w http.ResponseWriter, r *http.Request) {
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

func Start(router *mux.Router) {
	//This method takes the URL path and a function that will show the page
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/homepage", homePage).Methods("GET")
	router.HandleFunc("/feed", feedPage).Methods("GET")
	router.HandleFunc("/dashboard", dashboardPage).Methods("GET")
	router.HandleFunc("/dashboard/posts", dashboardPostsPage).Methods("GET")
	router.HandleFunc("/register", registerPage).Methods("GET")
	router.HandleFunc("/login", loginPage).Methods("GET")
	router.HandleFunc("/create-new-topic", createTopic).Methods("GET")
}
