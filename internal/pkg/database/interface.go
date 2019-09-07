package database

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type DBInterface interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Queryx(query string, args ...interface{}) (*sqlx.Rows, error)
	Ping() error
	Close() error
}
