# users
User management for Majero services.

## Running

```text
git clone https://github.com/majero-inc/users
cd users
go build
./users
```

Once the users binary is running you should be able to see the site at [http://localhost:8080](http://localhost:8080).

## DB Setup

### Sqlite
- make sure sqlite3 is installed
- setup users table (see below)

```
cat users_table.sql | sqlite3 majero.db
```

You can add users through the CLI, using `sqlite-add <first-name> <last-name>`

A list of users should be visible at [http://localhost:8080/sqlite](http://localhost:8080/sqlite).

## License
MIT
