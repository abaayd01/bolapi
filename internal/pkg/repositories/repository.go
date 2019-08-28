package repositories

import (
	"github.com/jmoiron/sqlx"
)

type BolAPIRepository struct {
	DB *sqlx.DB
}
