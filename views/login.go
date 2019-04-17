package views

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/majero-inc/users/modules/db"
)

func getLogin(w http.ResponseWriter, r *http.Request) {
	loginPage := Page{
		Title: "Login Page",
		Stylesheets: []string{
			"/public/css/base.css",
			"/public/css/login.css",
		},
		Data: map[string]string{
			"test": "",
		},
	}

	template, err := template.ParseFiles(
		"templates/base.html",
		"templates/login.html",
		"templates/header.html",
	)

	if err != nil {
		fmt.Println("Error Parsing File (login)")
		http.Redirect(w, r, "http://localhost:8080/home", http.StatusSeeOther)
	}

	template.Execute(w, loginPage)
}

func attemptLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.PostForm.Get("email"))
	fmt.Println(r.PostForm.Get("password"))

	if db.ValidateUser(r.PostForm.Get("email"), r.PostForm.Get("password")) {
		fmt.Print("\n\n Login Successful\n\n")
		http.Redirect(w, r, "http://localhost:8080/home", http.StatusSeeOther)
	} else {
		fmt.Print("\n\n Login Not Successful\n\n")
		http.Redirect(w, r, "http://localhost:8080/login", http.StatusSeeOther)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getLogin(w, r)
		break
	case "POST":
		r.ParseForm()
		attemptLogin(w, r)
		break
	default:
	}
}
