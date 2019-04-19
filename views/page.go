package views

type Page struct {
	Title       string
	Stylesheets []string
	Data        map[string][]string
	IsLoggedIn  bool
}
