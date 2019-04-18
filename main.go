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
	xList := []string{""}
	reader := bufio.NewReader(os.Stdin)
	printCLIHelper()

	for a > 0 {
		fmt.Print("Enter text: ")
		x, _ = reader.ReadString('\n')
		xList = strings.Split(strings.TrimRight(x, "\n"), " ")
		switch xList[0] {
		case "mysql-addDummy", "mysql-d":
				if db.CreateUser("test@test.com", "test") {
					fmt.Println("username: test@test.com, password: test. was added to the database\n")
				} else {
					fmt.Println("username: test@test.com, password: test. Already exist in the database\n")
				}
			break
		case "mysql-add", "mysql-a":
			if db.CreateUser(xList[1], xList[2]) {
				fmt.Println(fmt.Sprintf("username=%s, passaword=%s was added\n", xList[1], xList[2]))
			} else {
				fmt.Println(fmt.Sprintf("username=%s, passaword=%s already exists\n", xList[1], xList[2]))
			}				
			break
		case "mysql-help", "mysql-h", "mysql":
			fmt.Println("========================")
			fmt.Println("'mysql-help'    | 'mysql-h'  to get this text")
			fmt.Println("'mysql-addDummy'| 'mysql-d'  to add dummy object to database (test@test.com, test)")
			fmt.Println("'mysql-add'     | 'mysql-a'  [username] [password] to add user to database")
			fmt.Println("'mysql-print'   | 'mysql-p'  to print database")
			fmt.Println("========================\n")
			break
		case "mysql-print", "mysql-p":
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
