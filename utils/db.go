package utils

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	*sql.DB
	env Env
}

func NewDatabase(env Env) Database {
	if env.Environment == "production" {

		db, err := sql.Open("mysql", env.DatabaseUrl)
		if err != nil {
			panic(err)
		}
		return Database{db, env}
	} else if env.Environment == "test" {
		db, err := sql.Open("sqlite3", "./kplus.db")
		if err != nil {
			panic(err)
		}
		return Database{db, env}
	} else {

		panic("invalid environment")
	}
}
