package database

import "database/sql"

type Database struct {
	*sql.DB
}

func InitDB(db *sql.DB) *Database {
	return &Database{db}
}
