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
