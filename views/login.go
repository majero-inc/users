package views

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/majero-inc/users/modules/db"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func getLogin(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
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

		session.Values["authenticated"] = true
		session.Save(r, w)

		template.Execute(w, loginPage)
	} else {
		http.Error(w, "Already logged in", http.StatusForbidden)
		return
	}
}

func attemptLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
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
		attemptLogin(w, r)
		break
	default:
	}
}
