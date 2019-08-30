package database

import (
	"database/sql"
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

func (s *stub) Ping() error {
	return nil
}

func (s *stub) Close() error {
	return nil
}
