package sqlite

import (
    "log"
)

type User struct {
    FirstName string `db:"first_name"`
    LastName string `db:"last_name"`
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
        LastName: last,
    })

    conn.Close()
}
