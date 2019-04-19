package sqlite

import (
    "log"

    "upper.io/db.v3"
    "upper.io/db.v3/sqlite"
)

func Connect() db.Database {
    conn, err := sqlite.Open(Settings())
    if err != nil {
        log.Fatal(err)
    }

    return conn
}
