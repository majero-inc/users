package sqlite

import (
    "upper.io/db.v3/sqlite"
)

func Settings() sqlite.ConnectionURL {
    settings := sqlite.ConnectionURL {
        Database: "majero.db",
    }

    return settings
}

