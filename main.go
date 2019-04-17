package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/majero-inc/users/web"

	_ "github.com/go-sql-driver/mysql"
)

type USER struct {
	username string
	password string
	userID   sql.NullInt64
	groupID  sql.NullInt64
}

func mysql() {
	fmt.Print("Go connecting to database\n\n")

	db, err := sql.Open("mysql", "Maedz:test@tcp(127.0.0.1:3306)/demo")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	fmt.Print("Successfully connected to MySQL database\n\n")

	results, err := db.Query("SELECT * FROM USERS")

	if err != nil {
		panic(err.Error())
	}

	defer results.Close()

	for results.Next() {
		var user USER

		err = results.Scan(&user.username, &user.password, &user.userID, &user.groupID)
		if err != nil {
			panic(err.Error())
		}

		fmt.Print(user)
	}
	fmt.Print("\n\n")
}

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
		fmt.Println(x)
		switch strings.TrimRight(x, "\n") {
		case "mysql":
			mysql()
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
