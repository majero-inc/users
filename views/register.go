package views

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/majero-inc/users/modules/db"
)

func getRegister(w http.ResponseWriter, r *http.Request) {
	registerPage := Page{
		Title: "Register Page",
		Stylesheets: []string{
			"/public/css/base.css",
			"/public/css/login.css",
		},
	}

	template, err := template.ParseFiles(
		"templates/base.html",
		"templates/register.html",
		"templates/header.html",
	)

	if err != nil {
		fmt.Println("Error Parsing File (register)")
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}

	template.Execute(w, registerPage)
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.PostForm.Get("email"))
	fmt.Println(r.PostForm.Get("password"))

	if db.CreateUser(r.PostForm.Get("email"), r.PostForm.Get("password")) {
		fmt.Print("\n\n Successfully created user\n\n")
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	} else {
		fmt.Print("\n\n User was not created \n\n")
		http.Redirect(w, r, "/register", http.StatusSeeOther)
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getRegister(w, r)
		break
	case "POST":
		registerUser(w, r)
		break
	default:
	}
}
