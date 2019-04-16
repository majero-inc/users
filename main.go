package main

import (
    "net/http"
    "log"

    "github.com/majero-inc/users/web"
)

func main() {
    fs := http.FileServer(http.Dir("public"))
    http.Handle("/public/", http.StripPrefix("/public/", fs))

    web.Urls()

    log.Fatal(http.ListenAndServe(":8080", nil))
}
