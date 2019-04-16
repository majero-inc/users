package web

import (
    "net/http"

    "github.com/majero-inc/users/views"
)

func Urls() {
    http.HandleFunc("/", views.Home)
}
