package views

import (
    "html/template"
    "net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
    home_page := Page {
        Title: "Home",
        Stylesheets: []string {
            "/public/css/base.css",
        },
        Data: map[string] string {
            "test": "Home Page.",
        },
    }

    template, _ := template.ParseFiles("templates/base.html")

    template.Execute(w, home_page)
}
