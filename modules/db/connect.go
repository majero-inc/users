package db

import (
    "fmt"

    "github.com/majero-inc/users/modules/db/sqlite"
)

type Backend int
const (
    Sqlite   Backend = 0
)

func Connect(b Backend) {
    switch b {
    case Sqlite:
        sqlite.Connect()
        break
    default:
        fmt.Println("Backend not supported.")
    }
}
