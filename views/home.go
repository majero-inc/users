package views

import (
	"fmt"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	isLoggedIn := false
	session, _ := store.Get(r, "cookie-name")
	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		isLoggedIn = false
		fmt.Println("isLoggedIn false")
	} else {
		isLoggedIn = true
		fmt.Println("isLoggedIn true")
	}
	home_page := Page{
		Title: "Home",
		Stylesheets: []string{
			"/public/css/base.css",
		},
		Data: map[string]string{
			"test": "Home Page.",
		},
		IsLoggedIn: isLoggedIn,
	}

	template, err := template.ParseFiles(
		"templates/base.html",
		"templates/home.html",
		"templates/header.html",
	)

	if err != nil {
		fmt.Println("Error Parsing File (home)")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}

	template.Execute(w, home_page)
}
