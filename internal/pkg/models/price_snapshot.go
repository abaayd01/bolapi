package models

import "time"

type PriceSnapshot struct {
	Id          int64     `db:"price_snapshot_id"`
	CreatedTime time.Time `db:"created_time"`
	Price       float64   `db:"price"`
}

type PriceSnapshots []PriceSnapshot

func (pS *PriceSnapshots) TransformToFloatSlice() []float32 {
	var floatSlice []float32

	for _, priceSnapshot := range *pS {
		floatSlice = append(floatSlice, float32(priceSnapshot.Price))
	}

	return floatSlice
}
