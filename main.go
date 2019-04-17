package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/majero-inc/users/modules/db"

	"github.com/majero-inc/users/web"

	_ "github.com/go-sql-driver/mysql"
)

func printCLIHelper() {
	fmt.Print("Welcome to Majero CLI\nThese are the different tools\n\n")
	fmt.Print("'mysql' Connects to the database to enable debugging\n")
	fmt.Print("\nTo close the CLI use: 'q', 'quit'\n\n")
}

func mainloop() {
	x := " "
	a := 1
	reader := bufio.NewReader(os.Stdin)
	printCLIHelper()

	for a > 0 {
		fmt.Print("Enter text: ")
		x, _ = reader.ReadString('\n')
		switch strings.TrimRight(x, "\n") {
		case "mysql":
			db.PrintUsers()
			break
		case "quit", "q":
			a = 0
			break
		default:
			printCLIHelper()
		}
	}
	fmt.Print("Closing CLI\n")
	os.Exit(3)
}

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	fmt.Print("Starting CLI\n\n")
	go mainloop()

	web.Urls()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
