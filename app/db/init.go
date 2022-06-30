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
		chatid VARCHAR,
		birthday VARCHAR
	);`

	create_friend_table_sql := `CREATE TABLE IF NOT EXISTS friend (
		id INTEGER PRIMARY KEY,
		name VARCHAR,
		birthday VARCHAR,
		chatid INTEGER,
		userid INTEGER
	);`

	create_chat_table_sql := `CREATE TABLE IF NOT EXISTS chat (
		id INTEGER PRIMARY KEY,
		type VARCHAR,
		title VARCHAR,
		username VARCHAR,
		firstname VARCHAR,
		lastname VARCHAR,
		ownerid INTEGER
	);`

	create_user_chat_table_sql := `CREATE TABLE IF NOT EXISTS userchat (
		id INTEGER PRIMARY KEY,
		userid INTEGER,
		chatid INTEGER,
		UNIQUE (userid, chatid)
	);`

	for _, table := range []string{
		create_user_table_sql,
		create_friend_table_sql,
		create_chat_table_sql,
		create_user_chat_table_sql,
	} {
		err := create_table(client, table)
		if err != nil {
			return err
		}
	}

	return nil
}
