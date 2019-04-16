package views

import (
    "html/template"
    "net/http"
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

    template, _ := template.ParseFiles(
        "templates/base.html",
        "templates/login.html",
    )

    template.Execute(w, login_page)
}
