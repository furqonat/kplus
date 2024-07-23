package utils

import "database/sql"

type Database struct {
	*sql.DB
	env Env
}

func NewDatabase(env Env) *Database {
	db, err := sql.Open("mysql", env.DatabaseUrl)
	if err != nil {
		panic(err)
	}
	return &Database{db, env}
}
