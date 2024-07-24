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

func NewDatabase(env Env, d *sql.DB) Database {
	return Database{d, env}
}

func NewSqlDB(env Env) *sql.DB {
	if env.Environment == "production" {
		db, err := sql.Open("mysql", env.DatabaseUrl)
		if err != nil {
			panic(err)
		}
		return db
	} else if env.Environment == "test" {
		db, err := sql.Open("sqlite3", "./kplus.db")
		if err != nil {
			panic(err)
		}
		return db
	} else {
		panic("invalid environment")
	}
}
