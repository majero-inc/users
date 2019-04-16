package views

import (
    "html/template"
    "net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
    home_page := Page {
        Title: "Login",
        Stylesheets: []string {
            "/public/css/base.css",
        },
        Data: map[string] string {
            "test": "Login Page.",
        },
    }

    template, _ := template.ParseFiles("templates/base.html")

    template.Execute(w, home_page)
}
