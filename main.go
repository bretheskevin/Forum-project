package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		//set the status code to 404
		w.WriteHeader(http.StatusNotFound)

		page, errorParseFiles := template.ParseFiles("public/not-found.html")
		if errorParseFiles != nil {
			log.Fatal(errorParseFiles)
		}

		errorExecuteTemplate := page.ExecuteTemplate(w, "error.html", "")
		if errorExecuteTemplate != nil {
			return
		}
	} else {
		page, err := template.ParseFiles("public/index.html")
		if err != nil {
			log.Fatal(err)
		}
		errorExecuteTemplate := page.ExecuteTemplate(w, "index.html", "")
		if errorExecuteTemplate != nil {
			return
		}
	}
}

func main() {
	port := ":8080"

	// allow the server to access to the files
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))

	//This method takes the URL path and a function that will show the page
	http.HandleFunc("/", homePage)

	//start the server and use fmt to print the errors
	fmt.Println("Listening on http://localhost" + port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
}
