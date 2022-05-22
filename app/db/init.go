package db

import (
	"log"
	"database/sql"
)


func create_table(client *sql.DB, create_table_sql string) error {
	_, err := Client.Exec(create_table_sql)
	if err != nil {
        log.Println("Error when trying to prepare statement during creating tables")
        log.Println(err)
        return err
    }

	return nil
}


func create_tables(client *sql.DB) error {
	create_user_table_sql := `CREATE TABLE IF NOT EXISTS user (
		id INTEGER PRIMARY KEY,
		name VARCHAR,
		tgusername VARCHAR,
		chatid VARCHAR
	);`

	create_friend_table_sql := `CREATE TABLE IF NOT EXISTS friend (
		id INTEGER PRIMARY KEY,
		name VARCHAR,
		birthday VARCHAR,
		chatid INTEGER,
		userid INTEGER
	);`

	create_user_link_table_sql := `CREATE TABLE IF NOT EXISTS link (
		url VARCHAR,
		friendid VARCHAR
	);`

	create_congratulation_table_sql := `CREATE TABLE IF NOT EXISTS congratulations (
		id INTEGER PRIMARY KEY,
		text VARCHAR
	);`

	for _, table := range []string{
		create_user_table_sql,
		create_friend_table_sql,
		create_user_link_table_sql,
		create_congratulation_table_sql,
	} {
		err := create_table(client, table)
		if err != nil {
			return err
		}
	}

	return nil
}
