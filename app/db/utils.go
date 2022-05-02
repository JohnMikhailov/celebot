package db

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)


var Client *sql.DB


func init() {
    var err error
    Client, err = sql.Open("sqlite3", "celebot.db")
    if err != nil {
        panic(err)
    }
    if err = Client.Ping(); err != nil {
        panic(err)
    }

    create_tables(Client)

    fmt.Println("Database is ready")
}
