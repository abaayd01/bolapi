package database

import "database/sql"

type DBInterface interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Ping() error
	Close() error
}
