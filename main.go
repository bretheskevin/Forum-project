package main

import (
	"fmt"
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

func main() {
	port := ":8080"

	// allow the server to access to the files
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("public/images"))))
	http.Handle("/scripts/", http.StripPrefix("/scripts/", http.FileServer(http.Dir("public/scripts"))))
	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("public/styles"))))

	//This method takes the URL path and a function that will show the page
	http.HandleFunc("/", homePage)
	http.HandleFunc("/homepage", homePage)
	http.HandleFunc("/feed", feedPage)
	http.HandleFunc("/admin/dashboard/", dashboardPage)
	http.HandleFunc("/admin/dashboard/posts", dashboardPostsPage)
	http.HandleFunc("/register", registerPage)
	http.HandleFunc("/login", loginPage)


	//start the server and use fmt to print the errors
	fmt.Println("Listening on http://localhost" + port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
}
