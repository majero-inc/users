package views

import (
	"html/template"
	"log"
	"net/http"

	"github.com/majero-inc/users/modules/db/sqlite"
)

func Sqlite(w http.ResponseWriter, r *http.Request) {
	sqlite_page := Page{
		Title: "Sqlite Test Page",
		Stylesheets: []string{
			"/public/css/base.css",
		},
		Data: map[string][]string{
			"content": []string{"Sqlite test page."},
		},
		IsLoggedIn: false,
	}

	users := sqlite.Users()
	if users != nil {
		for _, user := range users {
			sqlite_page.Data["users"] = append(sqlite_page.Data["users"], user.FirstName+" "+user.LastName)
		}
	}

	templates, err := template.ParseFiles(
		"templates/base.html",
		"templates/header.html",
		"templates/sqlite.html",
	)
	if err != nil {
		log.Fatal(err)
	}

	templates.Execute(w, sqlite_page)
}
