package web

import (
	"net/http"

	"github.com/majero-inc/users/views"
)

func Urls() {
	http.HandleFunc("/", views.Home)
	http.HandleFunc("/home", views.Home)
	http.HandleFunc("/login", views.Login)
	http.HandleFunc("/register", views.Register)
	http.HandleFunc("/logout", views.Logout)
	http.HandleFunc("/sqlite", views.Sqlite)
}
