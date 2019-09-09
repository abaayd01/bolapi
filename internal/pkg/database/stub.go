package database

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
)

var DBStub *stub

type stub struct{}

func init() {
	newStub := stub{}
	DBStub = &newStub
}

func (s *stub) Exec(query string, args ...interface{}) (sql.Result, error) {
	return nil, nil
}

func (s *stub) NamedExec(query string, arg interface{}) (sql.Result, error) {
	return nil, nil
}

func (s *stub) NamedQuery(query string, arg interface{}) (*sqlx.Rows, error) {
	return nil, nil
}

func (s *stub) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return nil, nil
}

func (s *stub) Queryx(query string, args ...interface{}) (*sqlx.Rows, error) {
	return nil, nil
}

func (s *stub) QueryRow(query string, args ...interface{}) *sql.Row {
	return nil
}

func (s *stub) Ping() error {
	return nil
}

func (s *stub) Close() error {
	return nil
}
