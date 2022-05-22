package db

import (
    "database/sql"
    "log"
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

    log.Println("Database is ready")
}
