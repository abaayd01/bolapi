package repositories

import (
	"bolapi/internal/pkg/database"
	"bolapi/internal/pkg/models"
	"log"
)

type PriceSnapshotRepository interface {
	Insert(priceSnapshot *models.PriceSnapshot) error
}

type priceSnapshotRepository struct {
	DB database.DBInterface
}

func NewPriceSnapshotRepository(DB database.DBInterface) *priceSnapshotRepository {
	return &priceSnapshotRepository{DB: DB}
}

func (p *priceSnapshotRepository) Insert(priceSnapshot *models.PriceSnapshot) error {
	query := `
		INSERT INTO price_snapshots (created_time, price) VALUES ($1, $2)
	`

	_, err := p.DB.Exec(query, priceSnapshot.CreatedTime, priceSnapshot.Price)

	if err != nil {
		log.Printf("Error inserting PriceSnapshot")
		return err
	}

	return nil
}
