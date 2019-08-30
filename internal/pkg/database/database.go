package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 54320
	user   = "postgres"
	dbname = "bol_db"
)

var DB *sqlx.DB

func InitDB() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)

	db, err := sqlx.Open("postgres", psqlInfo)
	DB = db

	err = DB.Ping()
	if err != nil {
		return err
	}

	return nil
}

func CloseDB() {
	_ = DB.Close()
}
