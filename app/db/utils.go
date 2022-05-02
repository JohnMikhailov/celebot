package db

import (
    "database/sql"
    "fmt"
    "log"
    _ "github.com/mattn/go-sqlite3"

    "github.com/meehighlov/celebot/app"
)


var Client *sql.DB


func init() {
    config := app.GetConfig()

    host := config.DBHOST
    username := config.DBUSERNAME
    password := config.DBPASSWORD
    schema := config.DBSCHEMA

    connInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
        host, username, password, schema)

    var err error
    Client, err = sql.Open("sqlite3", connInfo)
    if err != nil {
        panic(err)
    }
    if err = Client.Ping(); err != nil {
        panic(err)
    }

    log.Println("Database ready to accept connections")
}
