package sqlite

import (
	"fmt"
	"log"

	"upper.io/db.v3"
)

type User struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
}

func Users() []User {
	conn := Connect()
	users_col := conn.Collection("users")

	res := users_col.Find()

	var users []User
	err := res.All(&users)
	if err != nil {
		log.Fatal(err)
	}

	conn.Close()

	return users
}

func AddUser(first, last string) {
	conn := Connect()
	users_col := conn.Collection("users")

	users_col.Insert(User{
		FirstName: first,
		LastName:  last,
	})

	conn.Close()
}

func UserAttemptLogin(username, password string) (bool, string) {
	conn := Connect()

	user_table := conn.Collection("users")

	fmt.Println("\nInput User: {'" + username + "', '" + password + "'}")

	res := user_table.Find(db.Cond{"first_name": username})

	var users []User

	res.All(&users)
	err := res.All(&users)
	if err != nil {
		log.Fatal(err)
	}

	if len(users) <= 0 {
		fmt.Println("\nNo Matching Users")
		return false, "failed"
	}

	fmt.Println("\nUser fetched on " + username + " : {'" + users[0].FirstName + "', '" + users[0].LastName + "'}")

	conn.Close()

	if users[0].LastName != password {
		return false, "failed"
	} else {
		return true, users[0].FirstName
	}
}
