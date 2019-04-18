package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //Use the mysql driver to simplify sql communication
)

type USER struct { //USER struct that needs to be used to extract information from USERS table
	username string
	password string
	userID   sql.NullInt64
	groupID  sql.NullInt64
}

func PrintUsers() {
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

func ValidateUser(user string, pass string) bool {
	fmt.Print("Go connecting to database\n\n")

	db, err := sql.Open("mysql", "Maedz:test@tcp(127.0.0.1:3306)/demo")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	fmt.Print("Successfully connected to MySQL database\n\n")

	s := fmt.Sprintf("SELECT * FROM USERS WHERE username='%s' AND password='%s'", user, pass)
	results, err := db.Query(s)

	if err != nil {
		panic(err.Error())
	}

	defer results.Close()

	if results.Next() {
		var user USER
		err = results.Scan(&user.username, &user.password, &user.userID, &user.groupID)
		if err != nil {
			fmt.Print("Did Not Find User (0)")
			return false
		} else {
			fmt.Print("Did Find User")
			fmt.Print(user)
			return true
		}

	} else {
		fmt.Print("Did Not Find User (1)")
		return false
	}
}

func CreateUser(user string, pass string) bool {
	fmt.Print("Go connecting to database\n\n")
	db, err := sql.Open("mysql", "Maedz:test@tcp(127.0.0.1:3306)/demo")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	fmt.Print("Successfully connected to MySQL database\n\n")

	selectQuery := fmt.Sprintf("SELECT * FROM USERS WHERE username='%s'", user)
	results, err := db.Query(selectQuery)

	if err != nil {
		panic(err.Error())
	}
	defer results.Close()

	if results.Next() {
		fmt.Print("User already exists")
		return false
	} else {
		fmt.Print("Did Not Find User, adding into database")
		insertQuery := fmt.Sprintf("insert into users (username,password) values ('%s','%s');",user, pass)
		insertResults, err := db.Query(insertQuery)
		if err != nil {
			panic(err.Error())
		}
		defer insertResults.Close()
		return true
	}
}