package views

import (
    "html/template"
    "net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
  tmpl, _ := template.ParseFiles("public/html/login.html")
  tmpl.Execute(w, "Login")
}
