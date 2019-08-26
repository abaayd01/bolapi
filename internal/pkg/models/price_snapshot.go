package models

import "time"

type PriceSnapshot struct {
	Id          int64     `db:"price_snapshot_id"`
	CreatedTime time.Time `db:"created_time"`
	Price       float64   `db:"price"`
}
