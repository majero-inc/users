package views

import (
    "html/template"
    "net/http"
    "fmt"
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

    template, err := template.ParseFiles(
        "templates/base.html",
        "templates/home.html",
        "templates/header.html",
    )

    if err != nil {
      fmt.Println("Error Parsing File (home)")
      w.WriteHeader(http.StatusInternalServerError)
      w.Write([]byte("500 - Something bad happened!"))
      return;
    }

    template.Execute(w, home_page)
}
