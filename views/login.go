package views

import (
    "html/template"
    "net/http"
    "fmt"
)

func Login(w http.ResponseWriter, r *http.Request) {
    login_page := Page {
        Title: "Login Page",
        Stylesheets: []string {
            "/public/css/base.css",
            "/public/css/login.css",
        },
        Data: map[string] string {
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

    template.Execute(w, login_page)
}
